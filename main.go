package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli/v2"
)

type responseBody struct {
	Translations []struct {
		SourceLanguage string `json:"detected_source_language"`
		Text           string `json:"text"`
	} `json:"translations"`
}

func main() {
	app := &cli.App{
		Name:    "deepl cli",
		Usage:   "Translate any text via deepl api.",
		Version: "0.0.2",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "to",
				Aliases: []string{"t"},
				Value:   "EN-US",
				Usage:   "language translate to",
			},
			&cli.StringFlag{
				Name:    "from",
				Aliases: []string{"f"},
				Value:   "JA",
				Usage:   "language translate from",
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() > 1 {
				fmt.Println("Please pass only one argument")
				return nil
			}

			client := &http.Client{
				CheckRedirect: func(req *http.Request, via []*http.Request) error {
					return http.ErrUseLastResponse
				},
			}

			req, err := http.NewRequest("GET", "https://api-free.deepl.com/v2/translate", nil)
			if err != nil {
				panic(err)
			}

			q := req.URL.Query()
			q.Add("text", c.Args().Get(0))
			q.Add("source_lang", c.String("from"))
			q.Add("target_lang", c.String("to"))
			q.Add("auth_key", os.Getenv("DEEPL_AUTH_KEY"))
			req.URL.RawQuery = q.Encode()

			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != 200 {
				fmt.Println(resp.Status)
				return nil
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			var respBody responseBody
			if err := json.Unmarshal(body, &respBody); err != nil {
				log.Fatal(err)
			}
			fmt.Println(respBody.Translations[0].Text)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

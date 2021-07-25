package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

type responseBody struct {
	Translations []struct {
		SourceLanguage string `json:"detected_source_language"`
		Text           string `json:"text"`
	} `json:"translations"`
}

func main() {
	app := cli.NewApp()
	app.Name = "deepl cli"
	app.Usage = "Translate any text via deepl api."
	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) error {
		if len(c.Args()) > 1 {
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
		q.Add("source_lang", "JA")
		q.Add("target_lang", "EN-US")
		q.Add("auth_key", os.Getenv("DEEPL_AUTH_KEY"))
		req.URL.RawQuery = q.Encode()

		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

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
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
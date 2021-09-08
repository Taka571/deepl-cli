
# Usage
1. `$ git clone https://github.com/Taka571/deepl-cli.git`
2. `$ cd deepl-cli`
3. `$ export DEEPL_AUTH_KEY='your key'` 
  
```
$ go run main.go 翻訳
translation
```

OR

`deepl-cli $ go install`
```
$ deepl-cli 猫
cat
```

# Options
- `-from alias: -f` language translate from 
- `-to   alias: -t` language translate to

Example(Translate from French to Japanese)
```
$ deepl-cli -f FR -t JA Traduction
翻訳
```

Please see below.  
https://www.deepl.com/docs-api/translating-text/request/


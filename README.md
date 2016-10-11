## Instructions

- Visit https://takeout.google.com/settings/takeout
- Download Chrome data (only need BrowserHistory)
- Extract `Takeout/Chrome/BrowserHistory.json` from archive
- Run app!

```
$ go run main.go urls -f BrowserHistory.json
...
$ go run main.go wiki -f BrowserHistory.json
```

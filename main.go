package main

import (
	"flag"
	"log"
	"fmt"
	"github.com/jchavannes/chrome-history-parser/wikipedia"
	"github.com/jchavannes/chrome-history-parser/bhp"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Available commands:")
		for _, cmd := range []string{
			"wiki",
			"urls",
		} {
			fmt.Println(" - " + cmd)
		}
		os.Exit(0)
	}

	flags := flag.NewFlagSet("", flag.ExitOnError)

	var filename string
	flags.StringVar(&filename, "f", "", "Location of BrowserHistory.json.")

	flags.Parse(args[1:])

	if len(filename) == 0 {
		flags.PrintDefaults()
		os.Exit(0)
	}

	browserHistory, err := bhp.LoadFromFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	switch args[0] {
	case "urls":
		for i := 0; i < len(browserHistory); i++ {
			b := browserHistory[i]
			fmt.Println(b.Time().String() + " - " + b.Url)
		}
	case "wiki":
		for i := 0; i < len(browserHistory); i++ {
			b := browserHistory[i]
			if !wikipedia.IsWikipediaUrl(b.Url) {
				continue
			}
			fmt.Println(b.Time().String() + " - " + wikipedia.ArticleNameFromUrl(b.Url))
		}
	}
}

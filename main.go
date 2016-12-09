package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
)

var (
	Version  string
	Revision string
)

// flags
var (
	userID     string
	perPage    string
	configPath string
)

type App struct {
	QiitaAPI QiitaAPI
	Config   Config
}

type Result struct {
	UserID  string
	Message string
}

var app = App{}

func (app App) Fetch(userID string, c chan<- Result) {
	stocks, err := app.QiitaAPI.Stocks(userID)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(stocks) == 0 {
		return
	}

	message := "## " + userID + "\n"
	for _, stock := range stocks {
		message += "- " + stock.Description() + "\n"
	}
	c <- Result{userID, message}
}

func (app App) run() error {
	var userIDs []string

	if len(userID) > 0 {
		userIDs = []string{userID}
	} else {
		userIDs = app.Config.QiitaUserIDs()
	}

	sort.Strings(userIDs)

	c := make(chan Result)
	for _, id := range userIDs {
		go app.Fetch(id, c)
	}

	for i := 0; i < len(userIDs); i++ {
		result := <-c
		fmt.Println(result.Message)
	}

	close(c)
	return nil
}

func init() {
	version := flag.Bool("v", false, "prints current qiist version")
	flag.StringVar(&userID, "u", "", "Qiita user ID")
	flag.StringVar(&perPage, "n", "5", "number of page. defalt: 5")
	flag.StringVar(&configPath, "c", "/path/to/config.json", "Default: ./config.json")
	flag.Parse()

	if *version {
		fmt.Println("Version: " + Version)
		fmt.Println("Revision: " + Revision)
		os.Exit(0)
	}

	if len(configPath) == 0 {
		configPath = "config.json"
	}

	config, err := NewConfig(configPath)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	app.Config = config
	app.QiitaAPI = NewQiitaAPI(config)
}

func main() {
	if err := app.run(); err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
}

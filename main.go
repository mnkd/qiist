package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"sync"
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

var app = App{}

func (app App) Fetch(userID string, wg *sync.WaitGroup) {
	defer wg.Done()

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
	fmt.Println(message)
}

func (app App) run() error {
	var userIDs []string

	if len(userID) > 0 {
		userIDs = []string{userID}
	} else {
		userIDs = app.Config.QiitaUserIDs()
	}

	sort.Strings(userIDs)

	var wg sync.WaitGroup
	for _, id := range userIDs {
		wg.Add(1)
		go app.Fetch(id, &wg)
	}

	wg.Wait()
	return nil
}

func init() {
	version := flag.Bool("v", false, "prints current qiist version")
	flag.StringVar(&userID, "user_id", "", "Qiita user ID")
	flag.StringVar(&perPage, "per_page", "5", "Defalt: 5")
	flag.StringVar(&configPath, "config", "/path/to/config.json", "Default: ./config.json")
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

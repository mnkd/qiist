package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
)

const (
	ExitCodeOK int = iota
	ExitCodeError
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
	if err != nil || len(stocks) == 0 {
		c <- Result{userID, ""}
		return
	}

	message := "## " + userID + "\n"
	for _, stock := range stocks {
		message += "- " + stock.Description() + "\n"
	}
	c <- Result{userID, message}
}

func (app App) Run() int {
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
		fmt.Fprintln(os.Stdout, result.Message)
	}

	close(c)
	return ExitCodeOK
}

func init() {
	version := flag.Bool("v", false, "prints current qiist version")
	flag.StringVar(&userID, "u", "", "Qiita user ID")
	flag.StringVar(&perPage, "n", "5", "number of page. defalt: 5")
	flag.StringVar(&configPath, "c", "/path/to/config.json", "Default: ./config.json")
	flag.Parse()

	if *version {
		fmt.Fprintln(os.Stdout, "Version:", Version)
		fmt.Fprintln(os.Stdout, "Revision:", Revision)
		os.Exit(ExitCodeOK)
	}

	if len(configPath) == 0 {
		configPath = "config.json"
	}

	config, err := NewConfig(configPath)
	if err != nil {
		os.Exit(ExitCodeError)
	}

	app.Config = config
	app.QiitaAPI = NewQiitaAPI(config)
}

func main() {
	os.Exit(app.Run())
}

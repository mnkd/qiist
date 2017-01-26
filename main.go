package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	ExitCodeOK int = iota
	ExitCodeError
)

var (
	Version  string
	Revision string
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
	users := app.Config.Qiita.Users
	c := make(chan Result)
	for _, user := range users {
		go app.Fetch(user, c)
	}

	for i := 0; i < len(users); i++ {
		result := <-c
		fmt.Fprintln(os.Stdout, result.Message)
	}

	close(c)
	return ExitCodeOK
}

func init() {
	var version bool
	var path string
	flag.BoolVar(&version, "v", false, "prints current qiist version")
	flag.StringVar(&path, "c", "/path/to/config.json", "Default: ./config.json")
	flag.Parse()

	if version {
		fmt.Fprintln(os.Stdout, "Version:", Version)
		fmt.Fprintln(os.Stdout, "Revision:", Revision)
		os.Exit(ExitCodeOK)
	}

	if len(path) == 0 {
		path = "config.json"
	}

	config, err := NewConfig(path)
	if err != nil {
		os.Exit(ExitCodeError)
	}

	app.Config = config
	app.QiitaAPI = NewQiitaAPI(config)
}

func main() {
	os.Exit(app.Run())
}

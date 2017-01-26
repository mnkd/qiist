package main

import (
	"fmt"
	"os"
)

type App struct {
	QiitaAPI QiitaAPI
	Config   Config
}

type Result struct {
	UserID  string
	Message string
}

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

func NewApp(config Config) App {
	var app = App{}
	app.Config = config
	app.QiitaAPI = NewQiitaAPI(config)
	return app
}

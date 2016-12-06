package main

import (
	"flag"
	"fmt"
	"os"
)

// flags
var (
	userID string
	perPage string
	configPath string
)

type App struct {
	QiitaAPI QiitaAPI
	Config Config
}
var app = App{}

func (app App) run() error {
	message := ""
	var userIDs []string

	if len(userID) > 0 {
		userIDs = []string{ userID }
	} else {
		userIDs = app.Config.QiitaUserIDs()
	}

	for _, id := range userIDs {
		stocks, err := app.QiitaAPI.Stocks(id)
		if err != nil {
			return err
		}

		message += "# " + id + "\n"
		for _, stock := range stocks {
			message += "- " + stock.Description() + "\n"
		}
		message += "\n"
	}

	fmt.Println(message)
	return nil
}

func init() {
	flag.StringVar(&userID, "user_id", "", "Qiita user ID")
	flag.StringVar(&perPage, "per_page", "5", "Defalt: 5")
	flag.StringVar(&configPath, "config", "/path/to/config.json", "Default: ./config.json")
	flag.Parse()

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

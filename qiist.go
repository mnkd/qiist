package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
    "time"
)

var (
	userID string
	perPage string
)

// Note:
// Qiita API (v2)
// https://qiita.com/api/v2/docs#get-apiv2usersuser_idstocks
// 指定されたユーザがストックした投稿一覧を、ストックした日時の降順で返します。
//
// ストックした日付は取得できない。

type Stock struct {
	Title     string `json:"title"`
	Url       string `json:"url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (stock Stock) dateDescription() string {
    // "created_at": "2000-01-01T00:00:00+00:00",
    jst := time.FixedZone("Asia/Tokyo", 9*60*60)
    t, err := time.Parse(time.RFC3339, stock.CreatedAt)
    if err != nil {
        return ""
    }
    return t.In(jst).Format("2006-01-02")
}

func (stock Stock) description() string {
	return "[" + stock.Title + "](" + stock.Url + ") (" + stock.dateDescription() + ")"
}

type App struct {
	Config Config
}

func (app *App) stocks(userID string) ([]Stock, error) {
	// Prepare HTTP Request
	url := "https://qiita.com/api/v2/users/" + userID + "/stocks?page=1&per_page=" + perPage
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("Authorization", "Bearer " + app.Config.Qiita.AccessToken)

	var stocks []Stock

	// Fetch Request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Failure : ", err)
		return stocks, err
	}

	// Read Response Body
	responseBody, _ := ioutil.ReadAll(response.Body)

	// Decode JSON
	if err := json.Unmarshal(responseBody, &stocks); err != nil {
		fmt.Println("JSON Unmarshal error:", userID, err)
		return stocks, err
	}

	return stocks, nil
}

func (app *App) run() error {
	message := ""
	var userIDs []string

	if len(userID) > 0 {
		userIDs = []string{ userID }
	} else {
		userIDs = app.Config.Qiita.UserIDs()
	}

	for _, id := range userIDs {
		stocks, err := app.stocks(id)
		if err != nil {
			return err
		}

		message += "# " + id + "\n"
		for _, stock := range stocks {
			message += "- " + stock.description() + "\n"
		}
		message += "\n"
	}

	fmt.Println(message)
	return nil
}

var app = App{}

func init() {
	flag.StringVar(&userID, "user_id", "", "Qiita user ID")
	flag.StringVar(&perPage, "per_page", "5", "Defalt: 5")
	flag.Parse()

	config, err := NewConfig()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	app.Config = config
}

func main() {
	if err := app.run(); err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
}

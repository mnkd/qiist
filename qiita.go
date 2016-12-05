package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Note:
// Qiita API (v2)
// https://qiita.com/api/v2/docs#get-apiv2usersuser_idstocks
// 指定されたユーザがストックした投稿一覧を、ストックした日時の降順で返します。
//
// ストックした日付は取得できない。

type QiitaAPI struct {
    Domain string
    AccessToken string
}

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

func (stock Stock) Description() string {
	return "[" + stock.Title + "](" + stock.Url + ") (" + stock.dateDescription() + ")"
}

func (qiita QiitaAPI) Stocks(userID string) ([]Stock, error) {
	// Prepare HTTP Request
	url := "https://" + qiita.Domain + "/api/v2/users/" + userID + "/stocks?page=1&per_page=" + perPage
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("Authorization", "Bearer " + qiita.AccessToken)

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

func NewQiitaAPI(config Config) QiitaAPI {
    qiita := QiitaAPI{}
    qiita.Domain = config.Qiita.Domain
    qiita.AccessToken = config.Qiita.AccessToken
    return qiita
}
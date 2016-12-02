package main

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type Config struct {
	Qiita Qiita `json:"qiita"`
}

type Qiita struct {
	AccessToken string `json:"access_token"`
	Users       []User `json:"users"`
}
type User struct {
	Id string `json:"id"`
}

func (qiita *Qiita) UserIDs() []string {
	var ids []string
	for _, user := range qiita.Users {
		ids = append(ids, user.Id)
	}
	return ids
}

func NewConfig() (Config, error) {
	var config Config

	str, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("Could not read config.json. ", err)
		return config, err
	}

	if err := json.Unmarshal(str, &config); err != nil {
		fmt.Println("JSON Unmarshal Error:", err)
		return config, err
	}

	return config, nil
}

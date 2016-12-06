package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Qiita struct {
		Domain      string `json:"domain"`
		AccessToken string `json:"access_token"`
		Users       []struct {
			Id string `json:"id"`
		} `json:"users"`
	} `json:"qiita"`
}

func (config Config) QiitaUserIDs() []string {
	var ids []string
	for _, user := range config.Qiita.Users {
		ids = append(ids, user.Id)
	}
	return ids
}

func NewConfig(path string) (Config, error) {
	var config Config

	str, err := ioutil.ReadFile(path)
	if err != nil {
		return config, err
	}

	if err := json.Unmarshal(str, &config); err != nil {
		return config, err
	}

	return config, nil
}

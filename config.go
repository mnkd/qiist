package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Qiita struct {
		Domain      string   `json:"domain"`
		AccessToken string   `json:"access_token"`
		PerPage     int      `json:"per_page"`
		Users       []string `json:"users"`
	} `json:"qiita"`
}

func NewConfig(path string) (Config, error) {
	var config Config

	str, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Config: <error> read config file:", path)
		return config, err
	}

	if err := json.Unmarshal(str, &config); err != nil {
		fmt.Fprintln(os.Stderr, "Config: <error> json unmarshal:", err, path)
		return config, err
	}

	return config, nil
}

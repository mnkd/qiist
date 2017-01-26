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

var app = App{}

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

	app = NewApp(config)
}

func main() {
	os.Exit(app.Run())
}

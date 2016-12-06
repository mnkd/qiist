#!/bin/sh
go fmt *.go
go build -o qiist main.go qiita.go config.go

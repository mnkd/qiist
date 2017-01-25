NAME     := qiist
VERSION  := 0.2.0
REVISION := $(shell git rev-parse --short HEAD)
SRCS    := $(shell find . -type f -name '*.go')
LDFLAGS := -ldflags="-X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\""

bin/$(NAME): $(SRCS) format
	go build $(LDFLAGS) -o bin/$(NAME)

format:
	go fmt $(SRCS)

clean:
	rm -rf bin/*

.PHONY: format clean

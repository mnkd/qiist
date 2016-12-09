NAME     := qiist
VERSION  := 0.1.2
REVISION := $(shell git rev-parse --short HEAD)

SRCS    := $(shell find . -type f -name '*.go')
LDFLAGS := -ldflags="-X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\""

bin/$(NAME): $(SRCS)
	go build $(LDFLAGS) -o bin/$(NAME)

.PHONY: format
format:
	go fmt $(SRCS)

.PHONY: clean
clean:
	rm -rf bin/*

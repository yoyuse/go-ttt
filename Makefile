# 2021-01-24
# 2017-06-07

SRC	= go-ttt.go
REPO	= github.com/yoyuse/go-ttt
GO_TTT	= $(REPO)
TTT	= $(REPO)/ttt

help:
	@echo "usage: make [run|test|build|install]"

run: $(SRC)
	@go run $< <<< "kdjd;sjdkd;s"

test:
	go test $(TTT)
	go test $(GO_TTT)

build:
	go build $(TTT)
	go build $(GO_TTT)

install:
	go install $(TTT)
	go install $(GO_TTT)


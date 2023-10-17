default:

build:
	wails build

setup:
	go get -u github.com/hajimehoshi/go-mp3
	go get -u github.com/hajimehoshi/oto/v2

modupdate:
	go get -u
	go mod tidy

.PHONY: default build setup modupdate

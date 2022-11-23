all: run

build:
    docker build . -t serviceNow-connector --platform=linux/amd64

run:
    go mod tidy
    go run main.go
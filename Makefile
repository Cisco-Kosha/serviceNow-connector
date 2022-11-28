all: run

build:
    docker build . -t servicenow-connector --platform=linux/amd64

run:
    go mod tidy
    go run main.go
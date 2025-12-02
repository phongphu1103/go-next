APP_NAME = api
APP_PATH = ./cmd/api
BIN_DIR = ./bin

.PHONY: run build clean tidy test

run:
	go run $(APP_PATH)

build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) $(APP_PATH)

clean:
	rm -rf $(BIN_DIR)

tidy:
	go mod tidy

test:
	go test ./...

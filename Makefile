NAME = sber-check
BUILD_DIR = ./bin/$(NAME)
CMD_DIR = ./cmd/$(NAME)


build:
	go build -v -o $(BUILD_DIR) $(CMD_DIR)


prod:
	@if [ "$(filter windows,$(MAKECMDGOALS))" != "" ]; then \
		GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)-win-x86.exe -v $(CMD_DIR); \
	elif [ "$(filter macos,$(MAKECMDGOALS))" != "" ]; then \
		GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)-darwin-amd64 -v $(CMD_DIR); \
	elif [ "$(filter linux-386,$(MAKECMDGOALS))" != "" ]; then \
		GOOS=linux GOARCH=386 go build -o $(BUILD_DIR)-linux-386 -v .$(CMD_DIR); \
	else \
		GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)-linux-amd64 -v $(CMD_DIR); \
	fi


run:
	go run $(CMD_DIR)/main.go


clean:
	rm -f ./bin/*

.PHONY: build, run, clean, prod

.DEFAULT_GOAL := build

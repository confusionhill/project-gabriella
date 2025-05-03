APP_NAME := offline-fable
MAIN := cmd/*.go
BUILD_DIR := build

run-dev:
	go run $(MAIN)

build-release-macos-x64:
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)-macos-x64 $(MAIN)

build-release-macos-arm:
	GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DIR)/$(APP_NAME)-macos-arm $(MAIN)

build-release-win-x64:
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)-win-x64.exe $(MAIN)

build-release-win-arm:
	GOOS=windows GOARCH=arm64 go build -o $(BUILD_DIR)/$(APP_NAME)-win-arm.exe $(MAIN)

build-release-linux-x64:
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)-linux-x64 $(MAIN)

build-release-linux-arm:
	GOOS=linux GOARCH=arm64 go build -o $(BUILD_DIR)/$(APP_NAME)-linux-arm $(MAIN)

BINARY_NAME=db-proxy
BUILD_DIR=public

build:
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 ./cmd/main.go
	GOOS=linux GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 ./cmd/main.go
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 ./cmd/main.go
	GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 ./cmd/main.go
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe ./cmd/main.go

clean:
	rm -f $(BUILD_DIR)/$(BINARY_NAME)-*

run: build
	export PUBLIC_IP=0.0.0.0:5432 && export DB_HOST=your-db-endpoint:5432 && ./$(BUILD_DIR)/$(BINARY_NAME)-linux-amd64

BINARY=app

.PHONY: clean run test init domain fmt

start:
	go run main.go

vendor:
	go mod vendor

build:
	go build -o dist/app main.go && cp .env.example dist/.env

test:
	go test ./..

migrate:
	cd ./database/migrations/ && goose mysql "root:@/asset_management?parseTime=true" up
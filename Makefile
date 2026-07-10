include .env

export

DATABASE_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)

run:
	go run main.go

build:
	go build ./...

test:
	go test ./...

fmt:
	gofmt -w .

tidy:
	go mod tidy

migrate-up:
	migrate -path migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path migrations -database "$(DATABASE_URL)" down 1
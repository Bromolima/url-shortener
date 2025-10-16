test:
	@go test ./internal/http/handler
	@go test ./internal/service 

migrate:
	@go run database/migration/main.go

docker-up:
	@docker-compose -f infra/docker-compose.yml up -d

docker-down:
	@docker-compose -f infra/docker-compose.yml down

run:
	@go run main.go

build:
	@go build -o main.go

install:
	@go install github.com/golang/mock
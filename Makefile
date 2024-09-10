build:
	@go build -o ./bin ./cmd/app/main.go

dev:
	@go run ./cmd/app/main.go

clean:
	@rm -rf ./bin

start:
	@./bin/main

air: 
	@air -c .air.toml
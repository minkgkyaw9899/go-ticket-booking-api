build:
	@go build -o ./bin ./cmd/app/main.go

run:
	@go run ./cmd/app/main.go

clean:
	@rm -rf ./bin
include cmd/local_config.env


run:
	docker compose up

run-local: 
	go run ./cmd/main.go
	
test:
	go test ./...
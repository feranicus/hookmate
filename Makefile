.PHONY: run tidy

run:
	@echo ">> starting server..."
	go run ./cmd/hookmate/main.go

tidy:
	@echo ">> tidying go modules..."
	go mod tidy
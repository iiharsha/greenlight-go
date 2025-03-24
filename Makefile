build:
	go build -o bin/greenlight ./cmd/api

run: build
	./bin/greenlight

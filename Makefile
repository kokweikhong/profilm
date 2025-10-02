build:
	go build -o bin/profilm cmd/profilm/main.go
run: build
	./bin/profilm
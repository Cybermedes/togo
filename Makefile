run:
	go run main.go

build:
	@go build -o bin/togo

execute:
	go build -o bin/togo && ./bin/togo
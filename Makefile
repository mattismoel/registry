run: build
	@./bin/reg

build:
	@go build -o ./bin/reg .

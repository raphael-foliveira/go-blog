build:
	go build -o bin/main .

run: build
	bin/main

dev:
	air

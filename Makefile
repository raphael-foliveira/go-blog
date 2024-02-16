build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/main .

build-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/main.exe .

build-osx:
	GOOS=darwin GOARCH=amd64 go build -o bin/main.darwin .

build: build-linux build-windows build-osx

run: build
	bin/main

dev:
	air

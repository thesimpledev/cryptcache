run: 
    go run ./cmd/cli

test:
    go test -coverprofile=coverage.out ./...

build:
	rm -rf dist/*
	GOOS=windows GOARCH=amd64 go build -o dist/cryptcache.exe ./cmd/cli
	GOOS=linux GOARCH=amd64 go build -o dist/cryptcache ./cmd/cli
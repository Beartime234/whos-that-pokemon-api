.PHONY: test build clean deploy

build:
	dep ensure -v
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/start api/start/*.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/check api/check/*.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/leaderboard api/leaderboard/*.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/name api/name/*.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

test:
	go test ./whosthatpokemon/... -v
	go test ./api/... -v

deploy: test clean build
	sls deploy -s dev --verbose

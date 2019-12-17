.PHONY: test build clean deploy

build:
	dep ensure -v
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/api api/*.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

test:
	go test ./api/... -v

deploy: clean build
	sls deploy --verbose

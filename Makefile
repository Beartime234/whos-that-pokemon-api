.PHONY: test build clean deploy

build:
	dep ensure -v
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/api api/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

test:
	go test ./api_test/... -v

deploy: clean build
	sls deploy --verbose

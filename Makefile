.PHONY: test build clean deploy

build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/main api/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

test:
	go test ./api/... -v

deploy: clean build
	sls deploy --verbose

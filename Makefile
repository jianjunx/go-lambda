.PHONY: build clean deploy gomodgen run

build:
	export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/apis ./main.go

deploy: clean build
	sls deploy --verbose

clean:
	rm -rf ./bin ./vendor

run:
	air -d

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh

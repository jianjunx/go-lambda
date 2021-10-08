.PHONY: build clean deploy gomodgen

build:
	export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/apis ./main.go

deploy: clean build
	sls deploy --verbose

clean:
	rm -rf ./bin ./vendor

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh

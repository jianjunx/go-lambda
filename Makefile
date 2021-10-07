.PHONY: build clean deploy gomodgen

build:
	export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/apis ./main.go
	mkdir bin/config
	cp config/release.config.yaml bin/config

clean:
	rm -rf ./bin ./vendor

deploy: clean build
	sls deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh

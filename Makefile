.PHONY: build clean deploy gomodgen run

build:
	export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/main ./main.go

deploy: clean build
	sls deploy --verbose

publish: clean build
	cp -r views bin
	cp -r config bin
	cp -r public bin

clean:
	rm -rf ./bin ./vendor

run:
	air -d

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh

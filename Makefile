all: clean test build
check-quality: lint fmt cyclo vet

APP=demo
ALL_PACKAGES=$(shell go list ./...)
SOURCE_DIRS=$(shell go list ./... | cut -d "/" -f3 | uniq)
VERSION ?= $(shell git describe --always)

clean:
	rm -rf ./out
	GO111MODULE=on go mod tidy -v

test:
	git config --global url."git@bitbucket.org:".insteadOf "https://bitbucket.org/"
	git config --global url."git@github.com:".insteadOf "https://github.com/"
	GO111MODULE=on go test -v -coverpkg=./... -coverprofile coverage.out ./...

build:
	@echo Building "./out/${APP}"...
	@mkdir -p ./out
	GO111MODULE=on go build -o "./out/${APP}" ./

fmt:
	gofmt -l -s -w $(SOURCE_DIRS)

imports:
	go get -u golang.org/x/tools/cmd/goimports
	goimports -l -w -v $(SOURCE_DIRS)

cyclo:
	go get -u github.com/fzipp/gocyclo
	gocyclo -over 8 $(SOURCE_DIRS)

vet:
	GO111MODULE=on go vet ./...
#linting
lint:
	go get -u golang.org/x/lint/golint
	golint -set_exit_status ./...
#Race detector
race:
	GO111MODULE=on go test -race -short ./...
#config copy
copy-config:
	cp config.yaml.sample config.yaml

#  docker compose build
docker-compose-build:
	docker-compose build --build-arg ssh_prv_key="$(shell cat ~/.ssh/id_rsa)" --build-arg ssh_pub_key="$(shell cat ~/.ssh/id_rsa.pub)"

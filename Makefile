APP_NAME=cyclo-cross

all: bin

.PHONY: run
run: bin
	@docker run --rm -d -p 4242:80 --name=$(APP_NAME)-dev -e PORT=80 $(APP_NAME)

.PHONY: stop
stop:
	@docker stop $(APP_NAME)-dev

.PHONY: bin
bin: 
	@docker build . -t $(APP_NAME) --build-arg LIBRARY_SOURCE_FOLDER=./example-c

.PHONY: unit-test
build-test: 
	@docker build . --target unit-test -t $(APP_NAME)-test --build-arg LIBRARY_SOURCE_FOLDER=./example-c
unit-test: build-test
	@docker run --rm $(APP_NAME)-test

go.mod:
	@docker run --rm -v ${CURDIR}:/src -w /src golang:1.14.3-alpine go mod init example

go.sum: go.mod
	@docker run --rm -v ${CURDIR}:/src -w /src golang:1.14.3-alpine go mod tidy
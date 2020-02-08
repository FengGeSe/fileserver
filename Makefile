# info
VERSION=v0.0.1

PROJECT_NAME=$(shell basename "$(PWD)")
GO_VERSION=$(shell go version  | awk '{print $$3}')
BUILD_TIME=$(shell date +%FT%T%z)
OS_ARCH=$(shell go version  | awk '{print $$4}')
CGO_ENABLED=0
GO_BIN=$(GOBIN)

GO_MODULE=$(shell sed -n '/module/p'  go.mod | awk '{print $$2}')

## help: Help for this project
help: Makefile
	@echo "Usage:\n  make [command]"
	@echo
	@echo "Available Commands:"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	
## build: Compile the binary.
build:
	@go build -o $(PROJECT_NAME) -ldflags "$(LDFLAGS)"

## install: build and install.
install:
	@go build -o $(PROJECT_NAME) -ldflags "$(LDFLAGS)" 
	@mv $(PROJECT_NAME) $(GO_BIN)

## build-linux: Compile the linux binary.
build-linux:
	@CGO_ENABLED=$(CGO_ENABLED) GOOS=linux GOARCH=amd64 go build -o $(PROJECT_NAME) -ldflags "$(LDFLAGS)"

## run: Build and run
run: build
	@./$(PROJECT_NAME)

## clean: Clean build files.
clean:
	rm -f $(PROJECT_NAME)
	

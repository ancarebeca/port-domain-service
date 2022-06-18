GO111MODULES=on
APP?=portdomain
REGISTRY?=gcr.io/images
COMMIT_SHA=$(shell git rev-parse --short HEAD)
PKGS=$(shell go list ./... | grep -v /vendor/)

# -----------------------------------------------------------------
#				Main targets
# -----------------------------------------------------------------

## all: cleans and builds the project
.PHONY: all
all: clean build

.PHONY: build
## build: builds the application
build: format
	@go build -o ${APP} main.go

## format: formats all packages
format:
	@go fmt $(PKGS)

.PHONY: run
## run: runs go run main.go
run:
	go run -race main.go

.PHONY: clean
## clean: cleans the binary
clean:
	@go clean

.PHONY: test
## test: runs go test with default values
test:
	go test -v -count=1 -race ./...

.PHONY: lint
## lint: runs linter
lint:
	@golangci-lint run

# -----------------------------------------------------------------
#				Docker targets
# -----------------------------------------------------------------

.PHONY: docker-build
## docker-build: builds the portdomain docker image to registry
docker-build:
	docker build -t ${APP}:${COMMIT_SHA} .

.PHONY: docker-up
## docker-up: starts the program instances
docker-up:
	docker-compose -f docker-compose.yml up -d

.PHONY: docker-down
## docker-down: stops the program instances, their databases and remove the containers
docker-down:
	docker-compose -f docker-compose.yml down -d


# -----------------------------------------------------------------
#			 Help
# -----------------------------------------------------------------
.PHONY: help
## help: Prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |	sed -e 's/^/ /'


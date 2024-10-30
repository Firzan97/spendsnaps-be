install:
	@go get ./...
	@go mod vendor

build:
	@echo "Cleaning output folder ..."
	@rm -rf output
	@echo "Installing dependencies to vendor folder ..."
	@go mod vendor
	@echo "Compiling core ..."
	@go build -o output/core

build_image:
	@docker build -t core .

clean:
	@rm -rf output

test:
	@go mod vendor && go test ./... -run ''

run:
	@./output/core

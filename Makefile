install:
	@go get ./...
	@go mod vendor

build:
	@echo "Cleaning output folder ..."
	@rm -rf output
	@echo "Installing dependencies to vendor folder ..."
	@go mod vendor
	@echo "Compiling receipt-api ..."
	@go build -o output/receipt-api

build_image:
	@docker build -t receipt-api .

clean:
	@rm -rf output

test:
	@go mod vendor && go test ./... -run ''

run:
	@./output/receipt-api

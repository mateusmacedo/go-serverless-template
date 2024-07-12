.PHONY: build clean deploy deploy-function up down

TARGET ?= ./cmd/hello/main.go
FUNCTION ?= hello
STAGE ?= local

build:
	env CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bootstrap $(TARGET)

clean:
	rm -rf ./bootstrap

up:
	docker compose -f docker/compose.yaml up -d

down:
	docker compose -f docker/compose.yaml down --volumes

deploy: clean build
	AWS_ENDPOINT_URL=http://localhost:4566 AWS_SECRET_ACCESS_KEY=secret AWS_ACCESS_KEY_ID=key AWS_DEFAULT_REGION=us-east-1 sls deploy --verbose --stage $(STAGE)

deploy-function: clean build
	AWS_ENDPOINT_URL=http://localhost:4566 AWS_SECRET_ACCESS_KEY=secret AWS_ACCESS_KEY_ID=key AWS_DEFAULT_REGION=us-east-1 sls deploy function -f $(FUNCTION) --verbose --stage $(STAGE)
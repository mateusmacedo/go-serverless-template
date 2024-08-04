.PHONY: clean up down deploy

STAGE ?= local

clean:
	rm -rf ./.bin ./vendor

up:
	docker compose -f docker/compose.yaml up -d

down:
	docker compose -f docker/compose.yaml down --volumes

deploy: clean
	AWS_ENDPOINT_URL=http://127.0.0.1:4566 AWS_SECRET_ACCESS_KEY=secret AWS_ACCESS_KEY_ID=key AWS_DEFAULT_REGION=us-east-1 ./node_modules/.bin/sls deploy --verbose --stage $(STAGE)
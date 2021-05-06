.PHONY: help
help:	## show makefile help
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

local-build:	## build simple-golang-api locally
	go build

docker-build:	## build simple-golang-api wholly using Docker
	docker build --tag nextmetaphor/simple-golang-api:latest .
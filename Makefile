USER_ID := $(shell id -u)
GROUP_ID := $(shell id -g)
GO_MODULES_CACHE := "${HOME}/go/pkg"

docker-compose-up:
	USER_ID="${USER_ID}" GROUP_ID="${GROUP_ID}" GO_MODULES_CACHE=${GO_MODULES_CACHE} docker-compose up

docker-compose-up-build:
	USER_ID="${USER_ID}" GROUP_ID="${GROUP_ID}" GO_MODULES_CACHE=${GO_MODULES_CACHE} docker-compose up --build

build-prod:
	docker build -f docker/prod/go/Dockerfile -t golang-boilerplate:latest .

docker-compose-down:
	docker-compose down
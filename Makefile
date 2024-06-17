BUF_VERSION:=v1.32.2
SWAGGER_UI_VERSION:=v4.15.5
DOCKER_CMD ?= docker

run:
	go run cmd/server/main.go

generate:
	go run github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION) generate

lint:
	go run github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION) breaking --against 'https://github.com/zcking/omega-catalog.git#branch=main'

docker:
	${DOCKER_CMD} build -t omega-catalog .

docker/run:
	${DOCKER_CMD} run --rm -it -p 8080:8080 -p 8081:8081 omega-catalog

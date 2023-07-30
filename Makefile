ifneq (,$(wildcard ./.env))
    include .env
    export
endif
TEST_COMMAND= GIN_MODE=test go test -v -cover ./src/controller/... ./src/middleware/...

start:
	POSTGRES_HOST=localhost go run ./src/main.go
update:
	go get -u ./...
setEnv:
	cp .env.example .env
setDbPath:
	mkdir ./db/data
build:
	docker build -t ${SERVER_IMAGE}:${SERVER_IMAGE_VERSION} .
buildApp:
	docker build -t ${WEB_IMAGE}:${WEB_IMAGE_VERSION} ./app
upDB:
	docker compose up db migrate
upAll:
	docker compose up
down:
	docker compose down
test:
	$(TEST_COMMAND)
testAll:
	go clean -testcache && $(TEST_COMMAND)
lint:
	golangci-lint run ./src/... --timeout=10m
generateFlowChart:
	docker run --rm -v ${PWD}/flowChart:/home/node/data matthewfeickert/mermaid-cli:latest -i $(FLOW_CHART).mmd -o $(FLOW_CHART).svg
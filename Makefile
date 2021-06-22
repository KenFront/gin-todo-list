TEST_COMMAND= GIN_MODE=test go test -v -cover ./src/controller/...

start:
	POSTGRES_HOST=localhost go run ./src/main.go
update:
	go get -u ./...
setEnv:
	cp .env.example .env
setDatabaseData:
	mkdir ./db/data
build:
	docker build -t kenfront/gin-todolist:latest .
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
	golangci-lint run ./src/...
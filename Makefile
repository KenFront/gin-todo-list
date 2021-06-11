start:
	POSTGRES_HOST=localhost go run ./src/main.go
setEnv:
	cp .env.example .env
build:
	docker build -t kenfront/gin-todolist:latest .
upDB:
	docker compose -f ./docker-compose.dev.yml up 
upAll:
	docker compose -f ./docker-compose.yml up
down:
	docker compose down
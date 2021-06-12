start:
	POSTGRES_HOST=localhost go run ./src/main.go
setEnv:
	cp .env.example .env
build:
	docker build -t kenfront/gin-todolist:latest .
upDB:
	docker compose up db
upAll:
	docker compose up
down:
	docker compose down
start:
	go run ./src/main.go
setEnv:
	cp .env.example .env
runDb:
	docker compose up
stopDb:
	docker compose down
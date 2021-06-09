start:
	go run ./src/main.go
runDb:
	docker compose up
stopDb:
	docker compose down
DOCKER_DEV_CONFIG=./docker-compose.dev.yml
DOCKER_PRODUCTION_CONFIG=./docker-compose.yml

start:
	POSTGRES_HOST=localhost go run ./src/main.go
setEnv:
	cp .env.example .env
build:
	docker build -t kenfront/gin-todolist:latest .
upDev:
	docker compose -f $(DOCKER_DEV_CONFIG) up 
downDev:
	docker compose  -f $(DOCKER_DEV_CONFIG) down
upProduction:
	docker compose -f $(DOCKER_PRODUCTION_CONFIG) up
downProduction:
	docker compose -f $(DOCKER_PRODUCTION_CONFIG) down
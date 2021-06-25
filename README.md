# gin-todo-list

## Description

Basic gin for Todo list

## Environment

- Docker: v20.10.7
- Go: v1.16.5
- golangci-lint: v1.41.0

## Service Flow Chart
<div align="center">
  <img src="./flowChart/services.svg" width="100%" alt="service flow chart">
</div>

## Command

```bash
# Launch go server
make start

# Update go module required packages
make update

# Set runtime enviroment
make setEnv

# Set database data volume path
make setDatabaseData

# Build go server image
make build

# Launch database with migrate
make upDB

# Launch all services
make upAll

# Shut down adn remove all services
make down

# Run unit test with cache
make test

# Run unit test without cache
make testAll

# Run golangci-lint
make lint

# Generate services flow chart
make generateFlowChart FLOW_CHART=services
```

## Development

```bash
# run first time
make setDatabaseData

# run first time or enviroment changed
make setEnv

# Launch database with migrate
make upDB

# Launch go server
make start
```

## Deployment

```bash
# Build go server image
make build

# Launch all services
make upAll
```

## Postman

- domain: http://localhost
  For local testing

## Todolist

- diagram
- Cron
- Swagger
- System log(MongoDb)
- Reverse Proxy(Nginx or Traefik)
- UI(React)

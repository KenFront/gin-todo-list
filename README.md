# gin-todo-list

## Description

Basic gin for Todo list

## Environment

- Docker: v20.10.7
- Go: v1.16.5
- golangci-lint: v1.41.0

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
- Nginx(Docker-compose)
- UI(React)

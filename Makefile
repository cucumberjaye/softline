-include .env
export $(shell sed 's/=.*//' .env)

DOCKER_COMPOSE=docker-compose

up:
	$(DOCKER_COMPOSE) up -d

down:
	$(DOCKER_COMPOSE) down

test:
	go test ./...
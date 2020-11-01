DOCKER=docker-compose
COMPOSE_PATH=build/docker-compose.yml
DOCKER_UP=$(DOCKER) -f $(COMPOSE_PATH) up
DOCKER_DOWN=$(DOCKER) -f $(COMPOSE_PATH) down

all:
run:
	$(DOCKER_UP) -d --build
down:
	$(DOCKER_DOWN) --volumes

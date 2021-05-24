## HELP
PROJECT           = needys-output-producer
## Colors
COLOR_RESET       = $(shell tput sgr0)
COLOR_ERROR       = $(shell tput setaf 1)
COLOR_COMMENT     = $(shell tput setaf 3)
COLOR_TITLE_BLOCK = $(shell tput setab 4)

## display this help text
help:
	@printf "\n"
	@printf "${COLOR_TITLE_BLOCK}${PROJECT} Makefile${COLOR_RESET}\n"
	@printf "\n"
	@printf "${COLOR_COMMENT}Usage:${COLOR_RESET}\n"
	@printf " make build\n\n"
	@printf "${COLOR_COMMENT}Available targets:${COLOR_RESET}\n"
	@awk '/^[a-zA-Z\-_0-9@]+:/ { \
				helpLine = match(lastLine, /^## (.*)/); \
				helpCommand = substr($$1, 0, index($$1, ":")); \
				helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
				printf " ${COLOR_INFO}%-15s${COLOR_RESET} %s\n", helpCommand, helpMessage; \
		} \
		{ lastLine = $$0 }' $(MAKEFILE_LIST)
	@printf "\n"

## stack - start the entire stack in background, then follow logs
start:
	docker-compose up --build --detach
	docker-compose logs --follow

## stack - stop the entire stack
stop:
	docker-compose down

## stack - only start the api "needys-output-producer"
api-only:
	docker-compose up needys-output-producer

## stack - only start the sidecars backends (rabbitmq here)
sidecars-only:
	docker-compose up rabbitmq

## docker - build the needys-output-producer image
build:
	docker build -t needys-output-producer:latest .

## docker - enter into the needys-output-producer container
enter:
	docker-compose exec needys-output-producer /bin/sh

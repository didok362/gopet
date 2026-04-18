include .env
export

export PROJECT_ROOT=$(shell pwd)


env-up:
	@mkdir -p out/pgdata
	@docker compose up -d gopet-postgres

env-down:
	@docker compose down gopet-postgres

env-cleanup:
	@read -p "Clear all volume files [y/n]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down gopet-postgres gopet-socat && \
		sudo rm -rf out/pgdata && \
		echo "File was cleared"; \
	else\
		echo "Canceled.";\
	fi

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Seq is empty. use 'make migrate-create seq=init'"; \
		exit 1; \
	fi; \
	docker compose run --rm gopet-migrate \
		create\
		-ext sql\
		-dir /migrations \
		-seq "$(seq)"

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Action is empty. use 'make migrate-action action=up'" \
		exit 1; \
	fi; \
	docker compose run --rm gopet-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@gopet-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"

socat-up:
	@docker compose up -d gopet-socat

socat-down:
	@docker compose down gopet-socat

gopet-run:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
	export POSTGRES_HOST=localhost && \
	go run cmd/gopet/main.go

logs-cleanup:
	@read -p "Clear all log files [y/n]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down gopet-postgres gopet-socat && \
		sudo rm -rf out/logs && \
		echo "File was cleared"; \
	else\
		echo "Canceled.";\
	fi

gopet-deploy:
	@docker compose up -d --build gopet

gopet-undeploy:
	@docker compose down gopet

ps:
	@docker compose ps

swagger-gen:
	@docker compose run --rm swagger \
		init \
		-g cmd/gopet/main.go \
		-o docs \
		--parseInternal \
		--parseDependency

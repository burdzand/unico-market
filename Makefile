.PHONY = up down build rebuild

up:
	chmod +x scripts/postgres/create-multiple-postgresql-databases.sh
	@docker-compose -f deployments/docker-compose.yml up -d  postgres
	echo Setting databases...
	sleep 5
	@docker-compose -f deployments/docker-compose.yml up -d --remove-orphans

build:
	@docker-compose -f deployments/docker-compose.yml build --no-cache

rebuild:
	@docker-compose -f deployments/docker-compose.yml  build --no-cache
	@docker-compose -f deployments/docker-compose.yml up -d --remove-orphans

down:
	@docker-compose -f deployments/docker-compose.yml  down --remove-orphans
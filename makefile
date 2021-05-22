build:
	docker-compose build
up:
	docker-compose up -d
start:
	@make build
	@make up
down:
	docker-compose down

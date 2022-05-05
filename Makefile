build:
	docker-compose build

run:
	docker-compose up


swagger:
	swag init -g cmd/main.go

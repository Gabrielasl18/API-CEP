run: 
	docker compose run --service-ports --rm go-api-cep

build:
	docker compose build go-api-cep --progress plain
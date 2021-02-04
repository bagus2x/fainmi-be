.PHONY: postgres adminer migrate

include .env

dev:
	go run api/main.go
adminer:
	docker run --rm -ti --network host adminer
postgres:
	docker run --rm -ti --network host -e POSTGRES_PASSWORD=admin123 postgres
migrate:
	migrate -source file://migrations -database ${DATABASE_URI} $(t)
testcache:
	go clean -testcache
coba:
	echo ${DEV_PORT}
.PHONY: postgres adminer migrate

export DEV_PORT=8080
export ACCESS_TOKEN_KEY=wndkub7t2et2y32yqknxlPsr9zigfsq
export REFRESH_TOKEN_KEY=nwdssqqswmdowjgknrgupuusqheie2
export DATABASE_URI=postgres://bagus2x:admin123@localhost:5432/fainmi

dev:
	export PORT=8080 && go run api/app.go
build:
	go build -o bin/app api/app.go
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

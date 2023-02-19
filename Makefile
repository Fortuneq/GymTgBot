postgres:
	docker run -d \
		--name session-db \
		-p 5432:5432 \
		-e POSTGRES_USER=root \
		-e POSTGRES_PASSWORD=pass \
		postgres:15-alpine

createdb:
	docker exec -it \
		session-db createdb \
		--username=root \
		--owner=root session

dropdb:
	docker exec -it \
		session-db dropdb \
		--username=root \
		session

migrateup:
	migrate \
		-path migrations \
		-database "postgresql://root:pass@localhost:5432/session?sslmode=disable" \
		-verbose up

migratedown:
	migrate \
    	-path migrations \
    	-database "postgresql://root:pass@localhost:5432/session?sslmode=disable" \
    	-verbose down

run:
	go run cmd/app/main.go

mock:
	mockgen -source=$(i) -destination=$(d)

test:
	go test -v -cover `go list ./... | grep -v ./*/mock`


lint:
	golangci-lint run --config=./.golangci.yml

protoc:
	protoc -I. --go-grpc_out=. --go_out=. proto/$(file).proto
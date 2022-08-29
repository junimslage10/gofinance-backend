createdb:
	createdb --username=postgres --owner=postgres go_finance

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=root -d postgres:14-alpine 

migrateup:
	migrate -path db/migration -database "postgresql://postgres:root@host.docker.internal:5432/go_finance?sslmode=disable" -verbose up

migrationdrop:
	migrate -path db/migration -database "postgresql://postgres:root@host.docker.internal:5432/go_finance?sslmode=disable" -verbose down

test:
	echo "127.0.0.1 host.docker.internal" >> /etc/hosts
	go test -v -cover ./...

server:
	go run main.go

.PHONY: createdb postgres dropdb migrateup migrationdrop test
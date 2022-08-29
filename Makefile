instalsqlc:
	go install github.com/kyleconroy/sqlc/cmd/sqlc@v1.15.0

migrateup:
	migrate -path db/migration -database "postgresql://postgres:root@host.docker.internal:5432/go_finance?sslmode=disable" -verbose up

migrationdrop:
	migrate -path db/migration -database "postgresql://postgres:root@host.docker.internal:5432/go_finance?sslmode=disable" -verbose down

gomodinit:
	go mod init github.com/junimslage10/gofinance-backend

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: migrateup migrationdrop test
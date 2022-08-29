instalsqlc:
	go install github.com/kyleconroy/sqlc/cmd/sqlc@v1.15.0

migrateup:
	migrate -path db/migration -database "postgresql://postgres:root@localhost:5432/go_finance?sslmode=disable" -verbose up

# migrationdrop:
# 	migrate -path db/migration -database "postgresql://postgres:root@localhost:5432/go_finance?sslmode=disable" -verbose down

# test:
# 	go test -v -cover ./...

# server:
# 	go run main.go

.PHONY: instalsqlc migrateup
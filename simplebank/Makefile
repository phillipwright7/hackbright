createdb:
	createdb simple_bank

dropdb:
	dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql:///simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql:///simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: createdb dropdb migrateup migratedown sqlc test
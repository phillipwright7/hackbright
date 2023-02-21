createdb:
	createdb ratings

dropdb:
	dropdb ratings

migrateup:
	migrate -path db/migrations -database "postgresql:///ratings?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql:///ratings?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

.PHONY: createdb dropdb migrateup migratedown sqlc
migrateup:
	migrate -path ./db/migrations -database "postgresql:///pgadmin-migrations?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql:///pgadmin-migrations?sslmode=disable" -verbose down

.PHONY: migrateup migratedown

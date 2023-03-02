ID =

createdb:
	createdb newcarsales

dropdb:
	dropdb newcarsales

migrateup:
	migrate -path db/migration -database "postgresql:///newcarsales?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql:///newcarsales?sslmode=disable" -verbose down

sqlc:
	sqlc generate

runserver:
	go run ./server/main.go

createcar:
	curl -i -X GET localhost:8080/createcar -d '{"make": "Toyota", "model": "Camry", "price": "25599.00", "transmission": "auto", "trim_level": "XSE", "color": "white"}'

deletecar:
	curl -i -X DELETE localhost:8080/deletecar/7

getcars:
	curl -i -X GET localhost:8080/getcars

getcar:
	curl -i -X GET localhost:8080/getcar/1

.PHONY: createdb dropdb migrateup migratedown sqlc runserver createcar deletecar getcars getcar
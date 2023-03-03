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
	curl -i -X DELETE localhost:8080/deletecar/$(ID)

getcars:
	curl -i -X GET localhost:8080/getcars

getcar:
	curl -i -X GET localhost:8080/getcar/$(ID)

createowner:
	curl -i -X GET localhost:8080/createowner -d '{"first_name": "James", "last_name": "Zebedee", "phone_number": 2055553456, "email": "james_zebedee@email.com"}'

deleteowner:
	curl -i -X DELETE localhost:8080/deleteowner/$(ID)

getowners:
	curl -i -X GET localhost:8080/getowners

getowner:
	curl -i -X GET localhost:8080/getowner/$(ID)

createsale:
	curl -i -X GET localhost:8080/createsale -d '{"car_id":{ "Int32": 9,"Valid": true }, "owner_id":{ "Int32": 5,"Valid": true }}'

deletesale:
	curl -i -X DELETE localhost:8080/deletesale/$(ID)

getsales:
	curl -i -X GET localhost:8080/getsales

getsale:
	curl -i -X GET localhost:8080/getsale/$(ID)

.PHONY: createdb dropdb migrateup migratedown sqlc runserver createcar deletecar getcars getcar createowner deleteowner getowners getowner createsale deletesale getsales getsale
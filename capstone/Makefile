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
	curl -i -X POST localhost:8080/createcar -d '{"make": "Toyota", "model": "Camry", "price": "25599.00", "transmission": "auto", "trim_level": "XSE", "color": "white"}'

deletecar:
	curl -i -X DELETE localhost:8080/deletecar/$(ID)

getcars:
	curl -i -X GET localhost:8080/getcars

getcar:
	curl -i -X GET localhost:8080/getcar/$(ID)

updatecar:
	curl -i -X PUT localhost:8080/updatecar/$(ID) -d '{"car_id": $(ID), "make": "Toyota", "model": "Camry", "price": "25599.00", "transmission": "auto", "trim_level": "XSE", "color": "white"}'

createowner:
	curl -i -X POST localhost:8080/createowner -d '{"first_name": "James", "last_name": "Zebedee", "phone_number": 2055553456, "email": "james_zebedee@email.com"}'

deleteowner:
	curl -i -X DELETE localhost:8080/deleteowner/$(ID)

getowners:
	curl -i -X GET localhost:8080/getowners

getowner:
	curl -i -X GET localhost:8080/getowner/$(ID)

updateowner:
	curl -i -X PUT localhost:8080/updateowner/$(ID) -d '{"owner_id": $(ID), "first_name": "James", "last_name": "Zebedee", "phone_number": 2055553456, "email": "james_zebedee@email.com"}'

createsale:
	curl -i -X POST localhost:8080/createsale -d '{"car_id":{ "Int32": 9,"Valid": true }, "owner_id":{ "Int32": 5,"Valid": true }}'

deletesale:
	curl -i -X DELETE localhost:8080/deletesale/$(ID)

getsales:
	curl -i -X GET localhost:8080/getsales

getsale:
	curl -i -X GET localhost:8080/getsale/$(ID)

updatesale:
	curl -i -X PUT localhost:8080/updatesale/$(ID) -d '{"sale_id": $(ID), "car_id":{ "Int32": 9,"Valid": true }, "owner_id":{ "Int32": 5,"Valid": true }}'

.PHONY: createdb dropdb migrateup migratedown sqlc runserver createcar deletecar getcars getcar createowner deleteowner getowners getowner createsale deletesale getsales getsale
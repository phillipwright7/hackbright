package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
	db "github.com/phillipwright7/hackbright/capstone/db/sqlc"
)

const (
	port     = ":8080"
	dbDriver = "postgres"
	dbSource = "postgresql:///newcarsales?sslmode=disable"
)

var queries *db.Queries

func main() {
	http.HandleFunc("/getcar/", GetAllCarsHandler)
	http.HandleFunc("/getcars", GetAllCarsHandler)
	http.HandleFunc("/createcar", CreateCarHandler)
	http.HandleFunc("/deletecar/", DeleteCarHandler)

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func OpenDatabase(w http.ResponseWriter) error {
	database, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		return err
	}

	queries = db.New(database)

	return nil
}

func CreateCarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")

		var c db.CreateCarParams

		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			log.Println(err)
			w.WriteHeader(400)
			return
		}

		defer r.Body.Close()

		if err := OpenDatabase(w); err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		car, err := queries.CreateCar(context.Background(), c)
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		payload := car.Model + " added to database."

		if err := json.NewEncoder(w).Encode(payload); err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}
	}
}

func DeleteCarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		w.Header().Set("Content-Type", "application/json")

		id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/deletecar/"))
		if err != nil {
			log.Println(err)
			w.WriteHeader(400)
			return
		}

		defer r.Body.Close()

		if err := OpenDatabase(w); err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		car, err := queries.GetCarDetails(context.Background(), int32(id))
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		if err := queries.DeleteCar(context.Background(), int32(id)); err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		payload := car.Model + " deleted from database."

		if err := json.NewEncoder(w).Encode(payload); err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}
	}
}

func GetAllCarsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")

		if err := OpenDatabase(w); err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		cars, err := queries.GetAllCars(context.Background())
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
		}

		if err := json.NewEncoder(w).Encode(cars); err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}
	}
}

func GetCarDetailsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")

		id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/getcar/"))
		if err != nil {
			log.Println(err)
			w.WriteHeader(400)
			return
		}

		defer r.Body.Close()

		if err := OpenDatabase(w); err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		car, err := queries.GetCarDetails(context.Background(), int32(id))
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		if err := json.NewEncoder(w).Encode(car); err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}
	}
}

package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
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
	http.HandleFunc("/createcar", createCarHandler)
	http.HandleFunc("/deletecar/", deleteCarHandler)
	http.HandleFunc("/getcars", getAllCarsHandler)
	http.HandleFunc("/getcar/", getCarDetailsHandler)
	http.HandleFunc("/createowner", createOwnerHandler)
	http.HandleFunc("/deleteowner/", deleteOwnerHandler)
	http.HandleFunc("/getowners", getAllOwnersHandler)
	http.HandleFunc("/getowner/", getOwnerDetailsHandler)
	http.HandleFunc("/createsale", createSaleHandler)
	http.HandleFunc("/deletesale/", deleteSaleHandler)
	http.HandleFunc("/getsales", getAllSalesHandler)
	http.HandleFunc("/getsale/", getSaleDetailsHandler)

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func openDatabase(w http.ResponseWriter) error {
	database, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		return err
	}

	queries = db.New(database)

	return nil
}

func createCarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")

		var c db.CreateCarParams

		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			log.Println(err)
			w.WriteHeader(400)
			return
		}
		if c == (db.CreateCarParams{}) {
			log.Println(errors.New("error: nil struct"))
			w.WriteHeader(400)
			return
		}

		defer r.Body.Close()

		if err := openDatabase(w); err != nil {
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
	} else {
		w.WriteHeader(400)
	}
}

func deleteCarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		w.Header().Set("Content-Type", "application/json")

		id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/deletecar/"))
		if err != nil {
			log.Println(err)
			w.WriteHeader(400)
			return
		}

		defer r.Body.Close()

		if err := openDatabase(w); err != nil {
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
	} else {
		w.WriteHeader(400)
	}
}

func getAllCarsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")

		defer r.Body.Close()

		if err := openDatabase(w); err != nil {
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
	} else {
		w.WriteHeader(400)
	}
}

func getCarDetailsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")

		id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/getcar/"))
		if err != nil {
			log.Println(err)
			w.WriteHeader(400)
			return
		}

		defer r.Body.Close()

		if err := openDatabase(w); err != nil {
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
	} else {
		w.WriteHeader(400)
	}
}

func createOwnerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")

		var o db.CreateOwnerParams

		if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
			log.Println(err)
			w.WriteHeader(400)
			return
		}
		if o == (db.CreateOwnerParams{}) {
			log.Println(errors.New("error: nil struct"))
			w.WriteHeader(400)
			return
		}

		defer r.Body.Close()

		if err := openDatabase(w); err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		owner, err := queries.CreateOwner(context.Background(), o)
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		payload := owner.FirstName + " " + owner.LastName + " added to database."

		if err := json.NewEncoder(w).Encode(payload); err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}
	} else {
		w.WriteHeader(400)
	}
}

func deleteOwnerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		w.Header().Set("Content-Type", "application/json")

		id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/deleteowner/"))
		if err != nil {
			log.Println(err)
			w.WriteHeader(400)
			return
		}

		defer r.Body.Close()

		if err := openDatabase(w); err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		owner, err := queries.GetOwnerDetails(context.Background(), int32(id))
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		if err := queries.DeleteOwner(context.Background(), int32(id)); err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		payload := owner.FirstName + " " + owner.LastName + " deleted from database."

		if err := json.NewEncoder(w).Encode(payload); err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}
	} else {
		w.WriteHeader(400)
	}
}

func getAllOwnersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")

		defer r.Body.Close()

		if err := openDatabase(w); err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		owners, err := queries.GetAllOwners(context.Background())
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
		}

		if err := json.NewEncoder(w).Encode(owners); err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}
	} else {
		w.WriteHeader(400)
	}
}

func getOwnerDetailsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")

		id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/getowner/"))
		if err != nil {
			log.Println(err)
			w.WriteHeader(400)
			return
		}

		defer r.Body.Close()

		if err := openDatabase(w); err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		owner, err := queries.GetOwnerDetails(context.Background(), int32(id))
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		if err := json.NewEncoder(w).Encode(owner); err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}
	} else {
		w.WriteHeader(400)
	}
}

func createSaleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")

		var s db.CreateSaleParams

		if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
			log.Println(err)
			w.WriteHeader(400)
			return
		}
		if s == (db.CreateSaleParams{}) {
			log.Println(errors.New("error: nil struct"))
			w.WriteHeader(400)
			return
		}

		defer r.Body.Close()

		if err := openDatabase(w); err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		sale, err := queries.CreateSale(context.Background(), s)
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		payload := "Sale ID " + fmt.Sprint(sale.SaleID) + " added to database."

		if err := json.NewEncoder(w).Encode(payload); err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}
	} else {
		w.WriteHeader(400)
	}
}

func deleteSaleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		w.Header().Set("Content-Type", "application/json")

		id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/deletesale/"))
		if err != nil {
			log.Println(err)
			w.WriteHeader(400)
			return
		}

		defer r.Body.Close()

		if err := openDatabase(w); err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		sale, err := queries.GetSaleDetails(context.Background(), int32(id))
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		if err := queries.DeleteSale(context.Background(), int32(id)); err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		payload := "Sale ID " + fmt.Sprint(sale.SaleID) + " deleted from database."

		if err := json.NewEncoder(w).Encode(payload); err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}
	} else {
		w.WriteHeader(400)
	}
}

func getAllSalesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")

		defer r.Body.Close()

		if err := openDatabase(w); err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		sales, err := queries.GetAllSales(context.Background())
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
		}

		if err := json.NewEncoder(w).Encode(sales); err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}
	} else {
		w.WriteHeader(400)
	}
}

func getSaleDetailsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")

		id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/getsale/"))
		if err != nil {
			log.Println(err)
			w.WriteHeader(400)
			return
		}

		defer r.Body.Close()

		if err := openDatabase(w); err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		sale, err := queries.GetSaleDetails(context.Background(), int32(id))
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		if err := json.NewEncoder(w).Encode(sale); err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}
	} else {
		w.WriteHeader(400)
	}
}

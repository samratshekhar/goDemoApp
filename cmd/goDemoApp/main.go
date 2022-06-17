package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"goDemoApp/internal/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS products
(
    id SERIAL,
    name TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT products_pkey PRIMARY KEY (id)
)`

func ensureTableExists(db *sql.DB) {
	if _, err := db.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func Initialize(user, password, dbname string) *sql.DB {
	connectionString :=
		fmt.Sprintf("%s:%s@tcp(%s)/demo", user, password, dbname)

	var err error
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/home", HomeHandler)
	http.Handle("/", r)
	db := Initialize("root", "root", "localhost:3306")
	ensureTableExists(db)
	p := models.Product{Id: 1, Name: "p2", Price: "66.66"}
	p.DeleteProduct(db)
	log.Fatal(http.ListenAndServe(":8080", r))

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Print("Home")
}

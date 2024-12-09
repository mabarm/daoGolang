package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" //used for referencing import; sql: unknown driver "postgres" (forgotten import?)
)

func main() {

	fmt.Println("Hello World!")

	//postgres connection string, user-> postgres , pass -> secret, host  -> localhsot , gopgtest -> db name
	connStr := "postgres://postgres:secret@localhost:5432/gopgtest?sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	//after main function finishes its execution this will get invoked
	// defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	//To check if there is a connection, established health check
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	createProductTable(db)

	fmt.Println("Successfully connected to db")
}

// In real life, use migrations
func createProductTable(db *sql.DB) {
	/*
	   	Product table will have
	   - IP
	   - Name
	   - Price
	   - Available
	   - Date created
	*/

	//6 numbers, 2 of which can be decimal
	//id auto incrementable
	//query string representing the sql that we will execute
	query := `CREATE TABLE IF NOT EXISTS product (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		price NUMERIC(6, 2) NOT NULL,
		available BOOLEAN,
		created timestamp DEFAULT NOW()
		
	)`

	//Executes the query against the database( here postgresql)
	_, err := db.Exec(query)

	if err != nil {
		log.Fatal(err)
	}

}

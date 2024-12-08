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

	fmt.Println("Successfully connected to db")
}

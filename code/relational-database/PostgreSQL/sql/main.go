package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {

	str := "host=1.13.102.61   port=5432 user=postgres password=pgsql0922     dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", str)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	db.Close()
}

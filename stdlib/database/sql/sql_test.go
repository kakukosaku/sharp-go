package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Example() {
	pool, err := sql.Open("mysql", "username:passwd@tcp(localhost:3306)/dbName")
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	err = pool.Ping()
	if err != nil {
		log.Fatal(err)
	}
	// Output:
	//
}

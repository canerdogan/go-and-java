package main

import (
	"database/sql"
	"fmt"
	_ "github.com/bmizerany/pq"
	"net/http"
)

const (
	port = 8080
)

var (
	config = ReadConfig()
	db     = openDB()
)

func main() {
	http.HandleFunc("/authenticate", Authenticate) // auth_resource.go
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func openDB() *sql.DB {
	db, err := sql.Open("postgres",
		fmt.Sprintf("dbname=%s user=%s password=%s sslmode=disable", config.Db.Name, config.Db.User, config.Db.Password))
	if err != nil {
		panic(err)
	}
	return db
}

package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Employee struct {
	Id      string `json:"id"`
	Names   string `json:"name"`
	Jobrole string `json:"jobrole"`
	Grade   int    `json:"grade"`
}

func main() {
	db, err := sqlx.Connect("postgres", "dbname=db_learn sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	employees := []Employee{}
	db.Select(&employees, "SELECT * FROM employee")
	// jason, john := employees[0], employees[1]

	fmt.Println(employees)

}

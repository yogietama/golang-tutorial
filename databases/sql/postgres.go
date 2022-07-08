package main

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = ""
	password = ""
	dbname   = "db_learn"
)

type Employee struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Jobrole string `json:"jobrole"`
	Grade   int    `json:"grade"`
}

func main() {
	psqlInfo := fmt.Sprintf("postgres:/%s/%s:@%s/%s?sslmode=disable", user, password, host, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("error===", err)
	}

	defer db.Close()

	getData := "SELECT * FROM employee"
	rows, _ := db.Query(getData)

	var employees []Employee
	for rows.Next() {
		employee := Employee{}

		s := reflect.ValueOf(&employee).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}
		err := rows.Scan(columns...)
		if err != nil {
			log.Fatal(err)
		}

		employees = append(employees, employee)
	}

	fmt.Println(employees)
}

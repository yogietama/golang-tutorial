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
	Names   string `json:"name"`
	Jobrole string `json:"jobrole"`
	Grade   int    `json:"grade"`
}

func getAllData(db *sql.DB) {
	getData := "SELECT * FROM employee"
	rows, _ := db.Query(getData)

	var employees []Employee
	for rows.Next() {
		// fmt.Println("print here if there is data")
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

func getTitleAndTimeoutData(db *sql.DB) {
	value := ""
	getData := "SELECT title, max_timeout_request FROM posts"
	rows, _ := db.Query(getData)

	var timeout sql.NullString

	for rows.Next() {
		res := ""
		err := rows.Scan(&value, &timeout)
		if err == nil {
			if timeout.Valid {
				res = value + "?timeout=" + timeout.String
			} else {
				res = value
			}

			fmt.Println(res)
		} else {
			fmt.Println(err)
		}
	}

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

	// getAllData(db)
	getTitleAndTimeoutData(db)
	// max_timeout_request

}

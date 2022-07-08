package main

import (
	"database/sql"
	"fmt"

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
	id      string
	name    string
	jobrole string
	grade   int
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
	rows, err := db.Query(getData)

	var employees []Employee
	for rows.Next() {
		var emp Employee

		if err := rows.Scan(&emp.id, &emp.name, &emp.jobrole, &emp.grade); err != nil {
			// fmt.Println(err)
			break
		}

		employees = append(employees, emp)
	}

	fmt.Println(employees)
}

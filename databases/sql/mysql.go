package main

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

type userDB struct {
	Id       int    `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = ""
	dbname   = "db_learn"
)

func main() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "db_learn",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	// db, err := sql.Open("mysql", psqlInfo)
	if err != nil {
		fmt.Println(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("error===", err)
	}

	defer db.Close()

	getData := "SELECT * FROM user"
	rows, _ := db.Query(getData)

	var users []userDB
	for rows.Next() {
		var user userDB

		if err := rows.Scan(&user.Id, &user.Email, &user.Password); err != nil {
			// fmt.Println(err)'
			break
		}

		users = append(users, user)

	}

	fmt.Println(users)
}

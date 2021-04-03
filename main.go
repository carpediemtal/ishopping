package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Connect("mysql", "root:asdf;lkj@(localhost:3306)/ishopping")
	if err != nil {
		fmt.Println("WTF")
	}

	rows, err := db.Query("SELECT * FROM user ")

	// iterate over each row
	for rows.Next() {
		var uid int
		// note that city can be NULL, so we use the NullString type
		var username string
		var password string
		var t int
		err = rows.Scan(&uid, &username, &password, &t)
		fmt.Println(uid, username, password, t)
	}

	// check the error from rows
	err = rows.Err()
}

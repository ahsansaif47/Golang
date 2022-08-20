package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type User struct {
	Name string `json:"name"`
}

func main() {
	psql_conn := fmt.Sprint("host=localhost port=5432 user=postgres " +
		"password=postgres database=test sslmode=disable")
	db, err := sql.Open("postgres", psql_conn)
	if err != nil {
		panic(err.Error())
	} else {
		// Inserting into database..
		insert, err := db.Query("INSERT into users VALUES('ABDULLAH')")
		if err != nil {
			panic(err.Error())
		} else {
			fmt.Println("Data Entered Sucessfully")
		}
		insert.Close()

		// Printing records from database..
		results, err := db.Query("Select name from users")
		if err != nil {
			panic(err.Error())
		}

		// Looping over records..
		for results.Next() {
			var user User
			err := results.Scan(&user.Name)
			if err != nil {
				panic(err.Error())
			}
			fmt.Println(user.Name)
		}
	}
	db.Close()
}

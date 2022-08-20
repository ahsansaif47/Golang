package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database connection definition..
// var db *gorm.DB
// var err error

type User struct {
	/*
		Defining gorm.Model to let gorm know that
		this is  goinf to be modelled somewhere in database.
	*/
	gorm.Model
	Name  string
	Email string
}

/*
	Writing the initial migration function
	Taking in the user struct and create table
	SQL statements for us..
*/
func initialMigration() {
	dsn := "host=localhost port=5432 user=postgres" +
		" password=postgres database=test sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to Connect to Database..")
	}
	db.AutoMigrate(&User{})
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	dsn := "host=localhost port=5432 user=postgres" +
		" password=postgres database=test sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not establish connection")
	}

	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func newUser(w http.ResponseWriter, r *http.Request) {
	dsn := "host=localhost port=5432 user=postgres" +
		" password=postgres database=test sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{Name: name, Email: email})
	fmt.Fprint(w, "New user sucessfully created..")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	dsn := "host=localhost port=5432 user=postgres" +
		" password=postgres database=test sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	search := mux.Vars(r)
	name := search["name"]
	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	fmt.Fprint(w, "User sucessfully deleted")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	dsn := "host=localhost port=5432 user=postgres " +
		"password=postgres database=test sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var updateUser User
	db.Where("name = ?", name).Find(&updateUser)
	updateUser.Email = email
	db.Save(&updateUser)

	fmt.Fprint(w, "User sucessfully updated")
}

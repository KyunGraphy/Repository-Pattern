package main

import (
	"fmt"
	"go-postgres/driver"
	"go-postgres/repository/repoimpl"
	models "go-postgres/model"
)

//host, port, user, password, dbname
const (
	host = "localhost"
	port = "5432"
	user = "ryan"
	password = "123456"
	dbname = "code4func"
)

func main() {
	// Create new connection, return PostgresDB object
	db := driver.Connect(host, port, user, password, dbname)

	err := db.SQL.Ping()
	if err != nil {
		panic(err)
	}

	// Create new UserRepoImpl object
	// Return an interface -> Call the method
	userRepo := repoimpl.NewUserRepo(db.SQL)

	// Create new model objects
	uhp := models.User{
		ID: 1,
		Name: "Ung Hoang Phuc",
		Gender: "Male",
		Email: "uhp@gmail.com",
	}

	dt := models.User{
		ID: 2,
		Name: "Dan Truong",
		Gender: "Male",
		Email: "dt@gmail.com",
	}

	userRepo.Insert(uhp)
	userRepo.Insert(dt)

	users, _ := userRepo.Select()

	for i:= range users {
		fmt.Println(users[i])
	}
}








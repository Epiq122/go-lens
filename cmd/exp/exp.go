package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"os"
)

type User struct {
	Name string
	Age  int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Exodus Epiq",
		Age:  44,
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}

	// fmt.Errorf
	err = CreateUser()
	if err != nil {
		log.Println(err)
	}
	err = CreateOrg()
	if err != nil {
		log.Println(err)
	}

}

// fmt.Errorf

func Connect() error {
	return errors.New("connection failed")
}

func CreateUser() error {
	err := Connect()
	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}
	return nil
}
func CreateOrg() error {
	err := CreateUser()
	if err != nil {
		return fmt.Errorf("create org: %w", err)

	}
	return nil
}

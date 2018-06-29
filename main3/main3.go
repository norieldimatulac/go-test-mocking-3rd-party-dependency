package main

import (
	"fmt"
)

func main() {

	var employee Employee

	employee.Db = NewDatabase()

	fullname, err := employee.GetFullName("emp:003")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(fullname)
	}

}

type Employee struct {
	ID        string
	FirstName string
	LastName  string
	Db        Database
}

func (e Employee) GetFullName(id string) (string, error) {

	err := e.Db.InitializeDB()
	if err != nil {
		return "", err
	}

	defer e.Db.CloseDB()

	r, err := e.Db.GetByID(id)
	if err != nil {
		return "", err
	}

	return r["FirstName"] + " " + r["LastName"], nil

}

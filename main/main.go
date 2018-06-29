package main

import (
	"errors"
	"fmt"
)

func main() {

	employee := Employee{
		FirstName: "Bronzon",
		LastName:  "Yatayan",
	}

	fullname, err := employee.GetFullName()

	if err != nil {
		fmt.Println(err.Error)
	} else {
		fmt.Println(fullname)
	}

}

type Employee struct {
	FirstName string
	LastName  string
}

func (e Employee) GetFullName() (string, error) {

	if e.FirstName == "" || e.LastName == "" {
		return "", errors.New("Invalid Name")
	}

	return e.FirstName + ", " + e.LastName, nil

}

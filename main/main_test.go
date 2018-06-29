package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFullName(t *testing.T) {
	assert := assert.New(t)

	tables := []struct {
		firstName         string
		lastName          string
		getFullNameResult string
		err               error
	}{
		{"Bronzon", "Yatayan", "Bronzon Yatayan", nil},
		{"", "Yatayan", "", errors.New("Invalid Name")},
		{"Bronzon", "", "", errors.New("Invalid Name")},
	}
	for _, table := range tables {

		employee := Employee{
			FirstName: table.firstName,
			LastName:  table.lastName,
		}

		fullname, err := employee.GetFullName()

		assert.Equal(fullname, table.getFullNameResult, "they should be equal")
		assert.Equal(err, table.err, "they should be equal")

	}

}

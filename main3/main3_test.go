package main

import (
	"errors"
	"go-unit-test/main3/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFullName(t *testing.T) {

	mockDB := &mocks.Database{}
	mockDB.On("InitializeDB").Return(nil)
	mockDB.On("CloseDB").Return(nil)
	mockDB.On("GetByID", "testid").Return(map[string]string{"FirstName": "Bronzon", "LastName": "Yatayan"}, nil)

	e := Employee{}
	e.Db = mockDB

	r, err := e.GetFullName("testid")

	mockDB.AssertExpectations(t)
	assert.Equal(t, "Bronzon Yatayan", r)
	assert.Equal(t, nil, err)

}

func TestGetFullNameDBError(t *testing.T) {
	mockDB := &mocks.Database{}
	mockDB.On("InitializeDB").Return(errors.New("some error"))

	e := Employee{}
	e.Db = mockDB

	_, err := e.GetFullName("testid")

	mockDB.AssertExpectations(t)

	if assert.Error(t, err) {
		assert.Equal(t, errors.New("some error"), err)
	}
}

func TestGetFullNameGetByIDError(t *testing.T) {
	mockDB := &mocks.Database{}
	mockDB.On("InitializeDB").Return(nil)
	mockDB.On("CloseDB").Return(nil)
	mockDB.On("GetByID", "testid").Return(map[string]string{}, errors.New("some error"))

	e := Employee{}
	e.Db = mockDB

	_, err := e.GetFullName("testid")

	mockDB.AssertExpectations(t)

	if assert.Error(t, err) {
		assert.Equal(t, errors.New("some error"), err)
	}
}

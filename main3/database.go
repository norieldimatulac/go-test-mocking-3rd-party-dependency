package main

type Database interface {
	InitializeDB() error
	CloseDB() error
	GetByID(string) (map[string]string, error)
}

func NewDatabase() Database {
	r := &Mongo{}
	return r
}

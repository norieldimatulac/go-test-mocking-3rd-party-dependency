package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Mongo struct {
	client *mgo.Session
}

func (r *Mongo) InitializeDB() error {

	var err error

	r.client, err = mgo.Dial("localhost:27017")
	return err

}

func (r *Mongo) CloseDB() error {
	if r.client != nil {
		r.client.Close()
	}
	return nil
}

func (r *Mongo) GetByID(id string) (map[string]string, error) {

	data := make(map[string]string)

	e := r.client.DB("test").C("employee")
	selector := bson.M{"EmpID": id}
	err := e.Find(selector).One(&data)
	if err != nil {
		return data, err
	}
	return data, nil

}

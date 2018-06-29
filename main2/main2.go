package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {

	var employee Employee

	fullname, err := employee.GetFullName("emp:001")

	if err != nil {
		fmt.Println(err.Error)
	} else {
		fmt.Println(fullname)
	}

}

type Employee struct {
	ID        string
	FirstName string
	LastName  string
}

func (e Employee) GetFullName(id string) (string, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     "10.10.202.122:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	r, err := client.HGetAll(id).Result()
	if err != nil {
		return "", err
	}

	return r["FirstName"] + " " + r["LastName"], nil

}

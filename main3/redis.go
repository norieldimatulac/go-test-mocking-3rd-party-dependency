package main

import (
	"github.com/go-redis/redis"
)

type Redis struct {
	client *redis.Client
}

func (r *Redis) InitializeDB() error {

	r.client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := r.client.Ping().Result()

	return err

}

func (r *Redis) CloseDB() error {
	return r.client.Close()
}

func (r *Redis) GetByID(id string) (map[string]string, error) {
	return r.client.HGetAll(id).Result()
}

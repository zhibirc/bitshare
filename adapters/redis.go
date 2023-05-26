package adapters

import "github.com/redis/go-redis/v9"

var dbClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0, // default DB
})

type Redis struct {}

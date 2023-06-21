// Package adapters implements wrappers for different data storage engines.
// This package offers unified API which then can be used in database service.
package adapters

import (
	"github.com/redis/go-redis/v9"
)

type Redis struct{}

func (_ *Redis) GetConnection() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0, // default DB
	})
}

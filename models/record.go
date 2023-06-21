// Package models contains representation of data structure.
package models

import (
	"context"

	"github.com/zhibirc/bitshare/services"
)

var storage = services.Storage.Connect()

type Record struct {
	//id   int
	Key  string
	Data string
	Ttl  int
	Hits int
}

func (rc *Record) Create(ctx context.Context, key string, data string, ttl int) error { // TODO: think about return value
	return storage.Set(ctx, key, data, ttl).Err()
}

func (rc *Record) GetOne(ctx context.Context, key string) (Record, error) {
	return storage.Get(ctx, key).Result()
}

func (rc *Record) GetAll() []Record

func (rc *Record) Delete(key string) error

// Package db implements database service.
// This service offers a set of operations for manipulating with stored data.
// The db package is abstracted from specific database API by using corresponding adapters.
package services

import (
	"os"

	"github.com/zhibirc/bitshare/adapters"
)

var storageEngine = os.Getenv("STORAGE_ENGINE")

type Storage struct{}

func (_ *Storage) chooseStorage() {
	if storageEngine == "redis" {
		return adapters.Redis
	}
}

func (inst *Storage) Connect() {
	return inst.chooseStorage()
}

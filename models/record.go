// Package models contains representation of data structure.
package models

type Record struct {
	//id   int
	Key  string
	Data string
	Ttl  int
	Hits int
}

func (rc *Record) Create(key string, data string, ttl int) int

func (rc *Record) GetOne(key string) (Record, error)

func (rc *Record) GetAll() []Record

func (rc *Record) Delete(key string) error

package main

type KeyGenerator interface {
	GenerateKey(text string) string
	TransformKey(key string) string
}

func generateId() string {
	return "123"
}
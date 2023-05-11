package main

type KeyGenerator interface {
	GenerateKey(text string) string
	TransformKey(key string) string
}

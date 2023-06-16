package tools

type KeyGenerator interface {
	GenerateKey(text string) string
	TransformKey(key string) string
}

func GenerateId() string {
	return "123" // stub
}

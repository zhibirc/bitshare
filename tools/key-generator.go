package tools

type KeyGenerator interface {
	GenerateKey(text string) string
	TransformKey(key string) string
}

func GenerateId() string {
	// TODO: use base36/62 IDs: lower/upper case Latin letters (only consonants) and numbers, without O, 0, I, i, 1.
	return "123" // stub
}

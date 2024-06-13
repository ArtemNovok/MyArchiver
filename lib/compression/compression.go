package compression

type Encode interface {
	Encode(str string) []byte
}

type Decode interface {
	Decode([]byte) string
}

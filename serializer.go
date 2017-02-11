package serializer

type Serializer interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(d []byte, v interface{}) error
	Name() string
}

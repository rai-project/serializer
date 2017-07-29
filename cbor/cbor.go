package cbor

import (
	"bytes"

	"github.com/rai-project/serializer"
	"github.com/ugorji/go/codec"
)

type cborSerializer struct{}

var handler codec.CborHandle

func (cborSerializer) Marshal(v interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := codec.NewEncoder(buf, &handler)
	err := enc.Encode(v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (cborSerializer) Unmarshal(d []byte, v interface{}) error {
	buf := bytes.NewBuffer(d)
	enc := codec.NewDecoder(buf, &handler)
	return enc.Decode(v)
}

func (cborSerializer) Name() string {
	return "cbor"
}

func New() serializer.Serializer {
	return cborSerializer{}
}

func init() {
	serializer.Register(cborSerializer{})
}

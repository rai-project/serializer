package binc

import (
	"bytes"

	"github.com/rai-project/serializer"
	"github.com/ugorji/go/codec"
)

type bincSerializer struct{}

var handler codec.BincHandle

func (bincSerializer) Marshal(v interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := codec.NewEncoder(buf, &handler)
	err := enc.Encode(v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (bincSerializer) Unmarshal(d []byte, v interface{}) error {
	buf := bytes.NewBuffer(d)
	enc := codec.NewDecoder(buf, &handler)
	return enc.Decode(v)
}

func (bincSerializer) Name() string {
	return "binc"
}

func New() serializer.Serializer {
	return bincSerializer{}
}

func init() {
	serializer.Register(bincSerializer{})
}

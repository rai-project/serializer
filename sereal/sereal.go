package sereal

import (
	"github.com/Sereal/Sereal/Go/sereal"
	"github.com/rai-project/serializer"
)

const name = "sereal"

// Codec that encodes to and decodes using Sereal.
// The Sereal codec has some interesting features, one of them being
// serialization of object references, including circular references.
// See https://github.com/Sereal/Sereal

type serealCodec struct{}

func (c serealCodec) Marshal(v interface{}) ([]byte, error) {
	return sereal.Marshal(v)
}

func (c serealCodec) Unmarshal(b []byte, v interface{}) error {
	return sereal.Unmarshal(b, v)
}

func (c serealCodec) Name() string {
	return name
}

func New() serializer.Serializer {
	return serealCodec{}
}

func init() {
	serializer.Register(serealCodec{})
}

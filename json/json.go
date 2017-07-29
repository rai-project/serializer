package json

import (
	"github.com/json-iterator/go"

	"github.com/rai-project/serializer"
)

type jsonSerializer struct{}

func (jsonSerializer) Marshal(v interface{}) ([]byte, error) {
	return jsoniter.Marshal(v)
}

func (jsonSerializer) Unmarshal(d []byte, v interface{}) error {
	return jsoniter.Unmarshal(d, v)
}

func (jsonSerializer) Name() string {
	return "json"
}

func New() serializer.Serializer {
	return jsonSerializer{}
}

func init() {
	serializer.Register(jsonSerializer{})
}

package json

import (
	"encoding/json"

	"github.com/rai-project/serializer"
)

type jsonSerializer struct{}

func (jsonSerializer) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (jsonSerializer) Unmarshal(d []byte, v interface{}) error {
	return json.Unmarshal(d, v)
}

func (jsonSerializer) Name() string {
	return "json"
}

func New() serializer.Serializer {
	return jsonSerializer{}
}

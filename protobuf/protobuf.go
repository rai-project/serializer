package protobuf

import (
	"errors"

	"github.com/gogo/protobuf/proto"
	"github.com/rai-project/serializer"
	"github.com/rai-project/serializer/json"
)

const name = "protobuf"

// More details on Protocol Buffers https://github.com/golang/protobuf
var (
	errNotProtocolBufferMessage = errors.New("value isn't a Protocol Buffers Message")
)

type protobufCodec struct{}

// Encode value with protocol buffer.
// If type isn't a Protocol buffer Message, gob encoder will be used instead.
func (c protobufCodec) Marshal(v interface{}) ([]byte, error) {
	message, ok := v.(proto.Message)
	if !ok {
		// toBytes() may need to encode non-protobuf type, if that occurs use json
		return json.New().Marshal(v)
	}
	return proto.Marshal(message)
}

func (c protobufCodec) Unmarshal(b []byte, v interface{}) error {
	message, ok := v.(proto.Message)
	if !ok {
		// toBytes() may have encoded non-protobuf type, if that occurs use json
		return json.New().Unmarshal(b, v)
	}
	return proto.Unmarshal(b, message)
}

func (c protobufCodec) Name() string {
	return name
}

func New() serializer.Serializer {
	return protobufCodec{}
}

func init() {
	serializer.Register(protobufCodec{})
}

package json

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/rai-project/serializer"
)

type jsonSerializer struct{}

var std jsonSerializer

func init() {
	std = New().(jsonSerializer)
}

func Marshal(v interface{}) ([]byte, error) {
	return std.Marshal(v)
}

func Unmarshal(d []byte, v interface{}) error {
	return std.Unmarshal(d, v)
}

func (jsonSerializer) Marshal(v interface{}) ([]byte, error) {
	// if c, ok := v.(codec.Selfer); ok {
	// 	w := new(bytes.Buffer)
	// 	bw := bufio.NewWriter(w)
	// 	h := new(codec.JsonHandle)
	// 	enc := codec.NewEncoder(bw, h)
	// 	err := enc.Encode(c)
	// 	bw.Flush()
	// 	return w.Bytes(), err
	// }
	return jsoniter.Marshal(v)
}

func (jsonSerializer) Unmarshal(d []byte, v interface{}) error {
	// if _, ok := v.(codec.Selfer); ok {
	// 	h := new(codec.JsonHandle)
	// 	dec := codec.NewDecoderBytes(d, h)
	// 	return dec.Decode(v)
	// }
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

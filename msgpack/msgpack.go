package msgpack

import (
	"bytes"
	"reflect"

	"github.com/rai-project/serializer"
	"github.com/ugorji/go/codec"
)

type msgpackSerializer struct{}

var handler codec.MsgpackHandle

func (msgpackSerializer) Marshal(v interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := codec.NewEncoder(buf, &handler)
	err := enc.Encode(v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (msgpackSerializer) Unmarshal(d []byte, v interface{}) error {
	buf := bytes.NewBuffer(d)
	enc := codec.NewDecoder(buf, &handler)
	return enc.Decode(v)
}

func (msgpackSerializer) Name() string {
	return "msgpack"
}

func New() serializer.Serializer {
	return msgpackSerializer{}
}

func init() {
	handler.MapType = reflect.TypeOf(map[string]interface{}(nil))
	serializer.Register(msgpackSerializer{})
}

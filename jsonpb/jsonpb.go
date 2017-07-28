package jsonpb

import (
	"bytes"
	"errors"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"

	"github.com/rai-project/serializer"
)

type jsonpbSerializer struct {
	marshaler   jsonpb.Marshaler
	unmarshaler jsonpb.Unmarshaler
}

func (s jsonpbSerializer) Marshal(v interface{}) ([]byte, error) {
	message, ok := v.(proto.Message)
	if !ok {
		return nil, errors.New("jsonpb marshaler requires a protobuf message")
	}
	var buf bytes.Buffer
	if err := s.marshaler.Marshal(&buf, message); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (s jsonpbSerializer) Unmarshal(d []byte, v interface{}) error {
	message, ok := v.(proto.Message)
	if !ok {
		return errors.New("jsonpb unmarshaler requires a protobuf message")
	}
	buf := bytes.NewBuffer(d)
	return s.unmarshaler.Unmarshal(buf, message)
}

func (jsonpbSerializer) Name() string {
	return "jsonpb"
}

func New() serializer.Serializer {
	return jsonpbSerializer{}
}

func init() {
	serializer.AddSerializer(jsonpbSerializer{
		marshaler: jsonpb.Marshaler{
			Indent: "  ",
		},
		unmarshaler: jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		},
	})
}

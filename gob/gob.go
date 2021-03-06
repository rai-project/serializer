package gob

import (
	"bytes"
	"encoding/gob"

	"github.com/rai-project/serializer"
)

const name = "gob"

// Codec serializing objects using the gob package.
// See https://golang.org/pkg/encoding/gob/
var Codec = new(gobCodec)

type gobCodec int

func (c gobCodec) Marshal(v interface{}) ([]byte, error) {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	err := enc.Encode(v)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (c gobCodec) Unmarshal(b []byte, v interface{}) error {
	r := bytes.NewReader(b)
	dec := gob.NewDecoder(r)
	return dec.Decode(v)
}

func (c gobCodec) Name() string {
	return name
}

func New() serializer.Serializer {
	return gobCodec{}
}

func init() {
	serializer.Register(gobCodec{})
}

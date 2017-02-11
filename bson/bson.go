package bson

import (
	"gopkg.in/mgo.v2/bson"

	"github.com/rai-project/serializer"
)

type bsonSerializer struct{}

func (bsonSerializer) Marshal(v interface{}) ([]byte, error) {
	return bson.Marshal(v)
}

func (bsonSerializer) Unmarshal(d []byte, v interface{}) error {
	return bson.Unmarshal(d, v)
}

func (bsonSerializer) Name() string {
	return "bson"
}

func New() serializer.Serializer {
	return bsonSerializer{}
}

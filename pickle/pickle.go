package pickle

import (
	"bytes"

	"github.com/hydrogen18/stalecucumber"
	"github.com/rai-project/serializer"
)

const name = "pickle"

// Codec that encodes to and decodes using Python Pickle.
// http://www.hydrogen18.com/blog/reading-pickled-data-in-go.html

type pickleCodec struct{}

func (c pickleCodec) Marshal(v interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	pickler := stalecucumber.NewPickler(buf)
	_, err := pickler.Pickle(v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (c pickleCodec) Unmarshal(b []byte, v interface{}) error {
	buf := bytes.NewBuffer(b)
	k, err := stalecucumber.Unpickle(buf)
	if err != nil {
		return err
	}
	return stalecucumber.UnpackInto(v).From(k, err)
}

func (c pickleCodec) Name() string {
	return name
}

func New() serializer.Serializer {
	return pickleCodec{}
}

func init() {
	serializer.Register(pickleCodec{})
}

package serializer

import (
	"errors"
	"strings"

	"golang.org/x/sync/syncmap"
)

var serializers syncmap.Map

func FromName(s string) (Serializer, error) {
	s = strings.ToLower(s)
	val, ok := serializers.Load(s)
	if !ok {
		log.WithField("serializer", s).
			Warn("cannot find serializer")
		return nil, errors.New("cannot find serializer")
	}
	serializer, ok := val.(Serializer)
	if !ok {
		log.WithField("serializer", s).
			Warn("invalid serializer")
		return nil, errors.New("invalid serializer")
	}
	return serializer, nil
}

func AddSerializer(s Serializer) {
	serializers.Store(strings.ToLower(s.Name()), s)
}

func Serializers() []string {
	names := []string{}
	serializers.Range(func(key, _ interface{}) bool {
		if name, ok := key.(string); ok {
			names = append(names, name)
		}
		return true
	})
	return names
}

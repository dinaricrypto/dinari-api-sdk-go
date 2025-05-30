package apiquery

import (
	"github.com/dinaricrypto/dinari-api-sdk-go/packages/param"
	"reflect"
)

func (e *encoder) newRichFieldTypeEncoder(t reflect.Type) encoderFunc {
	f, _ := t.FieldByName("Value")
	enc := e.typeEncoder(f.Type)
	return func(key string, value reflect.Value) ([]Pair, error) {
		if fielder, ok := value.Interface().(param.Optional); ok && fielder.IsPresent() {
			return enc(key, value.FieldByIndex(f.Index))
		} else if ok && fielder.IsNull() {
			return []Pair{{key, "null"}}, nil
		}
		return nil, nil
	}
}

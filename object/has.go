package object

import (
	"reflect"
)

func Has[T any](key string) func(T) bool {
	return func(obj T) bool {
		objType := reflect.TypeOf(obj)
		keyVal := reflect.ValueOf(key)
		valVal := reflect.ValueOf(obj)
		// nolint: exhaustive
		switch objType.Kind() {
		case reflect.Map:
			value := valVal.MapIndex(keyVal)

			return value.IsValid()
		case reflect.Struct:
			for i := range objType.NumField() {
				field := objType.Field(i)
				if field.Name == key {
					return true
				}
			}

			return false
		default:
			return false
		}
	}
}

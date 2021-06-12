package util

import (
	"reflect"
)

func IsSameType(a interface{}, b interface{}) bool {
	if a != nil && b != nil {
		typeA := reflect.TypeOf(a)
		typeB := reflect.TypeOf(b)
		return typeA.String() == typeB.String()
	}
	return a == nil && b == nil
}

package entities

import (
	"reflect"
)

type entity interface {
	Validation(allowNil bool) error
}

func validate(valueIsNil bool, allowNil bool, condition func() bool) (invalid bool) {
	if valueIsNil && !allowNil {
		return true
	}
	if valueIsNil && allowNil {
		return false
	}
	return condition()
}

func isNil(v interface{}) bool {
	return reflect.ValueOf(v).IsNil()
}

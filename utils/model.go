package utils

import "reflect"

func IsNotNil(a interface{}) bool {
	return !IsNil(a)
}

func IsNil(a interface{}) bool {
	defer func() { recover() }()
	return a == nil || reflect.ValueOf(a).IsNil()
}

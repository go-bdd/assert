package assert

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
)

func NotEquals(expected, actual interface{}) error {
	if err := validateEqualArgs(expected, actual); err != nil {
		return err
	}

	if ObjectsAreEqual(expected, actual) == true {
		return errors.New("the two objects are equals")
	}

	return nil
}

func Equals(expected, actual interface{}) error {
	if err := validateEqualArgs(expected, actual); err != nil {
		return err
	}

	if ObjectsAreEqual(expected, actual) == true {
		return nil
	}

	return fmt.Errorf("expected %+v but %+v given", expected, actual)
}

// ObjectsAreEqual determines if two objects are considered equal.
//
// This function does no assertion of any kind.
func ObjectsAreEqual(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	exp, ok := expected.([]byte)
	if !ok {
		return reflect.DeepEqual(expected, actual)
	}

	act, ok := actual.([]byte)
	if !ok {
		return false
	}
	if exp == nil || act == nil {
		return exp == nil && act == nil
	}
	return bytes.Equal(exp, act)
}

func isFunction(arg interface{}) bool {
	if arg == nil {
		return false
	}
	return reflect.TypeOf(arg).Kind() == reflect.Func
}

func validateEqualArgs(expected, actual interface{}) error {
	if isFunction(expected) || isFunction(actual) {
		return errors.New("cannot take func type as argument")
	}
	return nil
}

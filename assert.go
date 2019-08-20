package assert

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
)

// Determines if two objects are not equal.
func NotEquals(expected, actual interface{}) error {
	if err := validateEqualArgs(expected, actual); err != nil {
		return err
	}

	if ObjectsAreEqual(expected, actual) == true {
		return errors.New("the two objects are not equal")
	}

	return nil
}

// Determines if two objects are equal.
func Equals(expected, actual interface{}) error {
	if err := validateEqualArgs(expected, actual); err != nil {
		return err
	}

	if ObjectsAreEqual(expected, actual) == true {
		return nil
	}

	return fmt.Errorf("expected %+v but %+v given", expected, actual)
}

// Determines if two objects are considered equal.
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

func Nil(object interface{}) error  {
	if !isNil(object) {
		return fmt.Errorf("the object should be nil, but got %#v", object)
	}

	return nil
}

func NotNil(object interface{}) error  {
	if isNil(object) {
		return fmt.Errorf("the object should NOT be nil")
	}

	return nil
}

func isNil(object interface{}) bool {
	return object == nil
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

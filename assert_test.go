package assert_test

import "testing"
import "github.com/go-bdd/assert"

func TestEqual(t *testing.T) {
	if err := assert.Equals(1, 2); err == nil {
		t.Error(err)
	}

	if err := assert.Equals(1, func() {}); err == nil {
		t.Error("it shouldn't validate func in the param")
	}

	if err := assert.Equals(5, 5); err != nil {
		t.Error(err)
	}
}

func TestNotEqual(t *testing.T) {
	if err := assert.NotEquals(1, 2); err != nil {
		t.Error(err)
	}

	if err := assert.NotEquals(1, func() {}); err == nil {
		t.Error("it shouldn't validate func in the param")
	}

	if err := assert.NotEquals(5, 5); err == nil {
		t.Error(err)
	}
}

func TestNotNil(t *testing.T) {
	if err := assert.Nil(nil); err != nil {
		t.Error(err)
	}

	if err := assert.NotNil(123); err != nil {
		t.Error(err)
	}
}

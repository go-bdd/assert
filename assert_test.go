package assert_test

import "testing"
import "github.com/go-bdd/assert"

func TestEqual(t *testing.T) {
	if err := assert.Equals(1, 2); err == nil {
		t.Errorf("considered numbers should not be equal: %s", err)
	}

	if err := assert.Equals(5, 5); err != nil {
		t.Errorf("considered numbers should be equal: %s", err)
	}
}

func TestNotEqual(t *testing.T) {
	if err := assert.NotEquals(1, 2); err != nil {
		t.Errorf("considered numbers should be not equal: %s", err)
	}

	if err := assert.NotEquals(5, 5); err == nil {
		t.Errorf("considered numbers should be equal: %s", err)
	}
}

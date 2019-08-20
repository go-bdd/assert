# Assert library [![Build Status](https://travis-ci.com/go-bdd/assert.svg?branch=master)](https://travis-ci.com/go-bdd/assert) [![Coverage Status](https://coveralls.io/repos/github/go-bdd/assert/badge.svg?branch=master)](https://coveralls.io/github/go-bdd/assert?branch=master)

The aim of this project is to provide very simple and flexible assert library with as less dependencies as possible.
What's the main difference comparision to other libraries is that it just return an error when the validation fails.
It doesn't provide any complex error reporting.

## Usage

Available assertions:

* `Equals`
* `NotEquals`
* `ObjectsAreEqual`
* `Nil`
* `NotNil`

Here's an example:

```go
package main_test

import "github.com/go-bdd/assert"

func TestEqual(t *testing.T) {
	if err := assert.Equals(1, 2); err == nil {
		t.Errorf("considered numbers should not be equal: %s", err)
	}

	if err := assert.Equals(5, 5); err != nil {
		t.Errorf("considered numbers should be equal: %s", err)
	}
}
```

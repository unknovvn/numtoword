package assert

import (
	"fmt"
	"testing"
)

func Equal[T comparable](t *testing.T, expected T, actual T) bool {
	if expected != actual {
		t.Errorf(fmt.Sprintf("Invalid operation: %#v == %#v", expected, actual))

		return false
	}

	return true
}

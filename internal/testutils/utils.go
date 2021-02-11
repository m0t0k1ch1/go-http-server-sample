package testutils

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Equal checks whether actual and expected are equal.
func Equal(t *testing.T, actual interface{}, expected interface{}) {
	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Errorf("diff: %s", diff)
	}
}

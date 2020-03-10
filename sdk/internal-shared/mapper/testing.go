package mapper

import (
	"github.com/mitchellh/go-testing-interface"
)

// TestFunc is a helper that creates a Func and fails the test if it
// fails to create the Func.
func TestFunc(t testing.T, f interface{}) *Func {
	t.Helper()

	mf, err := NewFunc(f)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	return mf
}

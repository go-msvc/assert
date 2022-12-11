package assert

import (
	"testing"
)

func Bool(t *testing.T, title string, exp, got bool) {
	if exp != got {
		t.Errorf("Expected: %v", exp)
		t.Errorf("Got:      %v", got)
		t.Fatalf("%s: got %v != expected %v", title, got, exp)
	}
}

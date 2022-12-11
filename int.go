package assert

import (
	"testing"
)

func Int(t *testing.T, title string, exp, got int) {
	if exp != got {
		t.Errorf("Expected: %d", exp)
		t.Errorf("Got:      %d", got)
		t.Fatalf("%s: got %v != expected %v", title, got, exp)
	}
}

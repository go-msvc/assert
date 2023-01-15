package assert

import (
	"testing"

	"github.com/go-msvc/logger"
)

func Bool(t *testing.T, title string, exp, got bool) {
	if exp != got {
		t.Errorf("Expected: %v", exp)
		t.Errorf("Got:      %v", got)
		t.Fatalf("%V: %s: got %v != expected %v", logger.GetCaller(2), title, got, exp)
	}
}

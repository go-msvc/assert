package assert

import (
	"testing"

	"github.com/go-msvc/logger"
)

func Int(t *testing.T, title string, exp, got int) {
	if exp != got {
		t.Errorf("Expected: %d", exp)
		t.Errorf("Got:      %d", got)
		t.Fatalf("%V: %s: got %v != expected %v", logger.GetCaller(2), title, got, exp)
	}
}

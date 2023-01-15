package assert

import (
	"testing"

	"github.com/go-msvc/logger"
)

func String(t *testing.T, title string, exp, got string) {
	if exp != got {
		t.Errorf("Expected: (len=%d) %s", len(exp), exp)
		t.Errorf("Got:      (len=%d) %s", len(got), got)
		if len(exp) != len(got) {
			t.Fatalf("%V: %s got len %d != expected %d", logger.GetCaller(2), title, len(got), len(exp))
		}
		for i := 0; i < len(exp); i++ {
			if exp[i] != got[i] {
				t.Errorf("Got:      (len=%d) %s", len(got), got)
				t.Fatalf("%V: %s different at [%d]: got=\"%s\" != expected=\"%s\"", logger.GetCaller(2), title, i, got, exp)
			}
		}
		t.Fatalf("%V: %s not the same", logger.GetCaller(2), title) //not expected to get here...
	}
}

package assert

import (
	"testing"
)

func String(t *testing.T, title string, exp, got string) {
	if exp != got {
		t.Errorf("Expected: (len=%d) %s", len(exp), exp)
		t.Errorf("Got:      (len=%d) %s", len(got), got)
		if len(exp) != len(got) {
			t.Fatalf("%s got len %d != expected %d", title, len(got), len(exp))
		}
		for i := 0; i < len(exp); i++ {
			if exp[i] != got[i] {
				t.Errorf("Got:      (len=%d) %s", len(got), got)
				t.Fatalf("%s different at [%d]: got=\"%s\" != expected=\"%s\"", title, i, got, exp)
			}
		}
		t.Fatalf("%s not the same", title) //not expected to get here...
	}
}

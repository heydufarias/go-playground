package iteration

import "testing"

func TestRepet(t *testing.T) {
	repeted := Repeat("a")
	expect := "aaaaa"

	if repeted != expect {
		t.Errorf("expected %q but got %q", expect, repeted)
	}
}

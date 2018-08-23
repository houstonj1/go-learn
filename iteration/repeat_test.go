package iteration

import "testing"

// TestRepeat function
func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeated)
	}
}

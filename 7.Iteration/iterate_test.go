package main

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := 5 

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
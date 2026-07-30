package main

import (
	"testing"
)


func TestPerimeter(t *testing.T) {
	reactangle := Reactangle{10.0,10.0}
	got := Perimeter(reactangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	reactangle := Reactangle{12.0,6.0}
	got := Area(reactangle)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

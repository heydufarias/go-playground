package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Eduardo")
	want := "Hello, Eduardo"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

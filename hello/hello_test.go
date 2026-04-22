package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying 'Hello, Param' when adding a string", func(t *testing.T) {
		got := Hello("Eduardo")
		want := "Hello, Eduardo"
		assertCorrectMessage(t, got, want)

	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)

	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	// testing.TB is an interface from Go’s testing package
	// It’s implemented by:
	// *testing.T (unit tests)
	// *testing.B (benchmarks)
	// Meaning: this helper works for both tests and benchmarks.

	t.Helper()
	// t.Helper() marks the function as a helper function.
	// Practical effect: If an error happens, Go reports the line where the function was called, not inside it.
	// Without it:
	// error at assertCorrectMessage.go:3
	// With it:
	// error at hello_test.go:10
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

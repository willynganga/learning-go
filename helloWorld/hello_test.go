package helloWorld

import "testing"

func TestHelloWithName(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Will", "")
		want := "Hello, Will"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hola, World' when language provided is Spanish", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Saying hello to people in Spanish", func(t *testing.T) {
		got := Hello("Will", "Spanish")
		want := "Hola, Will"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Saying hello to people in French", func(t *testing.T) {
		got := Hello("Will", "French")
		want := "Bonjour, Will"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

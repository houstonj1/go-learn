package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("in French", func(t *testing.T) {
		got := Hello("James", "French")
		want := "Bonjour, James"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("James", "Spanish")
		want := "Hola, James"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("James", "")
		want := "Hello, James"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

}

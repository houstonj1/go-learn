package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("James")
	want := "Hello, James"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}

}

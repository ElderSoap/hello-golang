package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello Go"
	got := hello()
	if got != want {
		t.Errorf("want %s, got %s\n", want, got)
	}
}

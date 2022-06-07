package main

import "testing"

func hello() string {
	return "Hello Golang"
}

func TestHello(t *testing.T) {

	want := "Hello Golang"

	got := hello()

	if want != got {
		t.Fatalf("want %s, got %s\n", want, got)
	}
}

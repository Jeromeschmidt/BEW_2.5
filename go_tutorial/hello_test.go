package main

import "testing"

func TestHello(t *testing.T) {
    got := Hello("Jerome")
    want := "Hello, Jerome"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}

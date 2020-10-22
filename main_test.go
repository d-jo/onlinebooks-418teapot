package main

import "testing"

func TestMain(t *testing.T) {
	want := "asdf"
	if got := Tst(); got != want {
		t.Errorf("bad")
	}
}

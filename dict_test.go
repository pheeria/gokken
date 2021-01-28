package gokken

import "testing"

func TestDict(t *testing.T) {
	dict := Dict{size: 10}
	dict.Init()
	dict.Set("first", "word")
	dict.Set("second", "palabra")

	got := dict.Get("second")
	want := "palabra"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

package main

import("testing")

func TestAddr(t *testing.T){
	got := addr(4,5)
	want := 9
	if got!=want{
		t.Errorf("Error, got %q want %q", got, want)
	}
}
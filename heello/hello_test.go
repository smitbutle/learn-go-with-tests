package main

import(
	"testing"
)

func assertCorrectMessage(t *testing.T, got , want string){
	t.Helper()
	if got!=want{
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello( t *testing.T){
	t.Run("saying hello to people", func(t *testing.T){
		got:=Hello("Chris", "")
		want:="Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Bonjour, Smit", func(t *testing.T){
		got:=Hello("Smit", "French")
		want:="Bonjour, Smit"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T){
		got:=Hello("", "English")
		want:="Hello, World"
		assertCorrectMessage(t, got, want)
	})
}
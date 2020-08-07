package main

import "testing"

func TestDictionary(t *testing.T) {

	t.Run("Checking word find", func(t *testing.T) {

		dictionary := map[string]string{"test": "this is just a test"}

		got := Search(dictionary, "test")
		want := "this is just a test"

		assertStrings(t, got, want)

	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

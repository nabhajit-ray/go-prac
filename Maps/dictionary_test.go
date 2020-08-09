package main

import "testing"

func TestDictionary(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("Checking word find", func(t *testing.T) {

		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)

	})
	t.Run("Checking unknown word find", func(t *testing.T) {

		//dictionary := Dictionary{"test": "this is just a test"}

		_, got := dictionary.Search("unknown_word")
		//want := "word not found"

		assertError(t, got, ErrorNotFound)

	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test2", "This is test2")
	want := "This is test2"
	got, err := dictionary.Search("test2")
	if err != nil {
		t.Fatal("Something weird happended,could not add")
	}
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}

func TestUpdate(t *testing.T) {

	dictionary := Dictionary{"test": "this is test"}
	dictionary.Update("test", "this is a test2")
	want := "this is a test2"
	got, err := dictionary.Search("test")
	if err != nil {
		t.Fatal("Unexpected! could not update")
	}
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func TestDelete(t *testing.T) {
	dictionary := Dictionary{"test": "this is test"}
	dictionary.Delete("test")
	_, got := dictionary.Search("test")
	assertError(t, got, ErrorNotFound)

}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertString(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unkown")
		want := ErrNotFound
		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("new word", func(t *testing.T) {
		dictionary.Add("other", "this is another test")
		got, err := dictionary.Search("other")
		want := "this is another test"
		if err != nil {
			t.Fatal(err.Error())
		}
		assertString(t, got, want)
	})
	t.Run("existing word", func(t *testing.T) {
		err := dictionary.Add("test", "something")
		if err == nil {
			t.Fatal("expected an error but didn't get one")
		}
	})
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func assertError(t testing.TB, err, want error) {
	if err != want {
		t.Errorf("got %q, wanted %q", err, want)
	}
}

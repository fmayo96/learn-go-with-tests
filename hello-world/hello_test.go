package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say hello to a given name.", func(t *testing.T) {
		got := Hello("Franco")
		want := "Hello Franco!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Say Hello World! if no input name.", func(t *testing.T) {
		got := Hello()
		want := "Hello World!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

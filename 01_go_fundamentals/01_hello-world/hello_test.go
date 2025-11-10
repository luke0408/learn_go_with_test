package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("in Eng", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Kor", func(t *testing.T) {
		got := Hello("Chris", "Korean")
		want := "안녕하세요, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spa", func(t *testing.T) {
		got := Hello("Chris", "Spanish")
		want := "Hola, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Fre", func(t *testing.T) {
		got := Hello("Chris", "French")
		want := "Bonjour, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

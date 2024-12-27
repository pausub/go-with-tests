package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	t.Run("known word", func(t *testing.T) {
		key := "test"
		got, _ := dictionary.Search(key)
		want := "this is a test"

		assertString(t, got, want, key)
	})

	t.Run("unknown word", func(t *testing.T) {
		key := "void"
		_, err := dictionary.Search(key)

		if err == nil {
			t.Fatal("Expected error, got result")
		}

		assertString(t, err.Error(), ErrNotFound.Error(), key)
	})
}

func TestAdd(t *testing.T) {
	key := "city"
	value := "Kleckas"

	dictionary := Dictionary{}
	dictionary.Add(key, value)

	assertDictionary(t, dictionary, key, value)

}

func assertDictionary(t testing.TB, dictionary Dictionary, key, want string) {
	got, err := dictionary.Search(key)
	if err != nil {
		t.Fatal("Expected word, but got error", err)
	}

	assertString(t, got, want, key)
}

func assertString(t testing.TB, got, want, key string) {
	t.Helper()

	if got != want {
		t.Errorf("got: %q, want: %q, key: %q", got, want, key)
	}
}

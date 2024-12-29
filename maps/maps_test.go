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

	t.Run("add new word", func(t *testing.T) {
		key := "city"
		value := "Kleckas"

		dictionary := Dictionary{}
		err := dictionary.Add(key, value)

		assertError(t, err, nil)
		assertDictionary(t, dictionary, key, value)
	})

	t.Run("add existing word", func(t *testing.T) {
		key := "city"
		value := "Kleckas"

		dictionary := Dictionary{key: value}

		err := dictionary.Add(key, value)

		assertError(t, err, ErrWordExists)
		assertDictionary(t, dictionary, key, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update existing word", func(t *testing.T) {
		key := "city"
		value := "Kleckas"

		dictionary := Dictionary{key: value}

		updatedValue := "Kaunas"

		err := dictionary.Update(key, updatedValue)

		assertError(t, err, nil)
		assertDictionary(t, dictionary, key, updatedValue)
	})

	t.Run("update non existing word", func(t *testing.T) {
		key := "city"
		dictionary := Dictionary{}

		value := "Kaunas"

		err := dictionary.Update(key, value)

		assertError(t, err, ErrWordDoesntExist)
		assertNotExists(t, dictionary, key)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete existing word", func(t *testing.T) {
		key := "city"
		value := "Kleckas"

		dictionary := Dictionary{key: value}

		err := dictionary.Delete(key)

		assertError(t, err, nil)
		assertNotExists(t, dictionary, key)
	})

	t.Run("delete non  existing word", func(t *testing.T) {
		key := "city"

		dictionary := Dictionary{}

		err := dictionary.Delete(key)

		assertError(t, err, ErrWordDoesntExist)
		assertNotExists(t, dictionary, key)
	})
}

func assertNotExists(t testing.TB, d Dictionary, key string) {
	t.Helper()

	_, exists := d[key]
	if exists {
		t.Errorf("key %q unexpectedly exists in dictionary", key)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}

func assertDictionary(t testing.TB, dictionary Dictionary, key, want string) {
	t.Helper()

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

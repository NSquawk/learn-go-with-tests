package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	t.Run("known key", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown key", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		assertErrors(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}

	t.Run("add new word", func(t *testing.T) {
		word := "add"
		definition := "test adding new word"
		dictionary.Add(word, definition)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("add existing word", func(t *testing.T) {
		word := "add"
		newDefinition := "redifine on add"
		definition := "test adding new word"

		err := dictionary.Add(word, newDefinition)
		assertErrors(t, err, ErrAlreadyExists)
		assertDefinition(t, dictionary, word, definition)

	})
}

func TestUpdate(t *testing.T) {
	dictionary := Dictionary{}

	t.Run("update definition", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		newDefinition := "testing an update"
		dictionary.Add(word, definition)
		dictionary.Update(word, newDefinition)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("update word that doesn't exist", func(t *testing.T) {
		word := "update"
		definition := "this is a test"
		err := dictionary.Update(word, definition)
		assertErrors(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	dictionary := Dictionary{}
	word := "test"
	definition := "this is a test"
	dictionary.Add(word, definition)
	dictionary.Delete(word)
	_, err := dictionary.Search(word)

	if err != ErrNotFound {
		t.Errorf("expected %q to be deleted", word)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q but wanted %q", got, want)
	}
}

func assertErrors(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q but wanted %q", got, want)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("expected word from disctionary but got error:", err)
	}

	assertStrings(t, got, definition)
}

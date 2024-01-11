package maps101

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"foo": "bar"}

	t.Run("map with custom type", func(t *testing.T) {
		got, _ := dictionary.Search("foo")
		want := "bar"
		assertStrings(t, got, want)
	})
	t.Run("search word that does not exist", func(t *testing.T) {
		_, err := dictionary.Search("buzz")
		assertError(t, err, ErrNotFound)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func TestAddWord(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "foo"
		definition := "bar"

		dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "foo"
		definition := "bar"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "buzz")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find word added word:", err)
	}

	assertStrings(t, got, definition)
}

func TestUpdate(t *testing.T) {
	t.Run("update existing word", func(t *testing.T) {
		word := "foo"
		definition := "bar"
		dictionary := Dictionary{word: definition}
		newDefinition := "buzz"

		dictionary.Update(word, newDefinition)

		assertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("update non existing word", func(t *testing.T) {
		word := "foo"
		definition := "bar"
		dictionary := Dictionary{word: definition}
		newDefinition := "buzz"

		err := dictionary.Update("fooo", newDefinition)

		assertError(t, err, ErrWordDoesNotExist)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestDelete(t *testing.T) {
	word := "foo"
	definition := "bar"
	dictionary := Dictionary{word: definition}

	dictionary.Delete(word)
	_, err := dictionary.Search(word)

	assertError(t, err, ErrNotFound)
}

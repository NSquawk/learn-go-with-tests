package maps

type Dictionary map[string]string

var (
	ErrNotFound         = DictionaryErr("word not found in the dictionary")
	ErrAlreadyExists    = DictionaryErr("word already exists in the dictionary")
	ErrWordDoesNotExist = DictionaryErr("word does not exist in the dictionary to be updated")
)

type DictionaryErr string

func (d DictionaryErr) Error() string {
	return string(d)
}

func (d Dictionary) Search(searchKey string) (string, error) {
	definition, ok := d[searchKey]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, ok := d[word]

	if ok {
		return ErrAlreadyExists
	}

	d[word] = definition
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
		return nil
	default:
		return err
	}
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

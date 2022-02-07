package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound      = errors.New("Not found")
	errDuplicateWord = errors.New("The word already exists")
	errCantDelete    = errors.New("Cant delete the word that not exists")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	// 값, 존재여부
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word into a dic
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		d[word] = def
		return nil
	} else {
		return errDuplicateWord
	}
}

// update a word if exists
func (d Dictionary) Update(word, def string) {
	d[word] = def
}

// delete a word if exists
func (d Dictionary) Delete(word string) (string, error) {
	_, err := d.Search(word)
	if err != nil {
		return "", errCantDelete
	} else {
		delete(d, word)
		return word + " is deleted", nil
	}
}

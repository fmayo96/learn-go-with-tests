package maps

import "errors"

var ErrNotFound = errors.New("word not in dictionary")
var ErrAlreadyExists = errors.New("word already in dictionary")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(key, value string) error {
	_, ok := d[key]
	if ok {
		return ErrAlreadyExists
	}
	d[key] = value
	return nil
}

package mydict

import "errors"

// Dictionary Type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errAlreadyExist = errors.New("Key is already used. Try to different key.")

func (d Dictionary) Search(word string) (string, error) {
	val, exist := d[word]
	if exist {
		return val, nil
	}
	return "", errNotFound
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case errNotFound:
		d[key] = value
	case nil:
		return errAlreadyExist
	}
	return nil
}

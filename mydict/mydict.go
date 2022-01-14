package mydict

import "errors"

// Dictionary Type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")

func (d Dictionary) Search(word string) (string, error) {
	val, exist := d[word]
	if exist {
		return val, nil
	}
	return "", errNotFound
}

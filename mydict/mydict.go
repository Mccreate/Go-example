package mydict

import (
	"errors"
	"fmt"
)

// Dictionary Type
type Dictionary map[string]string

var (
	errNotFound     = errors.New("Not Found.")
	errAlreadyExist = errors.New("Key is already used. Try to different key.")
	errCantUpdate   = errors.New("Can't update non-existing word.")
)

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
	fmt.Println("Key:", key+", Value:", value, "is Added.")
	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case nil:
		d[key] = value
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

func (d Dictionary) Delete(key string) {
	_, err := d.Search(key)
	if err != errNotFound {
		delete(d, key)
	}
}

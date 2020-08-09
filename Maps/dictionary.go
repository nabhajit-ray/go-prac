package main

import "errors"

type Dictionary map[string]string

var ErrorNotFound = errors.New("Word not found")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if ok != true {
		return "", ErrorNotFound

	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) {

	d[word] = definition
}

func (d Dictionary) Update(word, definition string) error {
	_, ok := d[word]
	if ok != true {
		return ErrorNotFound
	}
	d[word] = definition
	return nil

}

func (d Dictionary) Delete(word string) error {

	_, ok := d[word]
	if ok != true {
		return ErrorNotFound
	} else {
		delete(d, word)
		return nil
	}
}

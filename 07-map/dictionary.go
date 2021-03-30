package _7_map

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func (dict Dictionary)Search(key string) (string, error)  {
	value, ok := dict[key]
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}

func (dict Dictionary) Add(key, value string) {
	dict[key] = value
}

/*Version 1
func Search(dictionary Dictionary, key string) (string, error) {
	value, ok := dictionary[key]
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}

*/
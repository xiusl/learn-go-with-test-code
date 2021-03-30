package _7_map

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound = errors.New("could not find the word you were looking for")
	ErrWorkExists = errors.New("cannot add word because it already exists")
)

func (dict Dictionary)Search(key string) (string, error)  {
	value, ok := dict[key]
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}

func (dict Dictionary) Add(key, value string) error {
	_, err := dict.Search(key)

	switch err {
	case ErrNotFound:
		dict[key] = value
	case nil:
		return ErrWorkExists
	default:
		return err
	}
	return nil
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

/*NOTE
func (dict Dictionary) Add(key, value string) {}
	这里注意到，dict 是值传递而不是指针，因为 map 是引用类型，因此值传递的时候也可以修改他的值
	引用类型不会因为作为值传递产生拷贝，只会有一个副本
*/
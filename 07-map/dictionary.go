package _7_map

type Dictionary map[string]string

var (
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExists = DictionaryErr("cannot update word because it does not exists")
)

type DictionaryErr string

func (err DictionaryErr) Error() string {
	return string(err)
}

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
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (dict Dictionary) Update(key, value string) error {
	_, err := dict.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExists
	case nil:
		dict[key] = value
	default:
		return err
	}
	return nil
}

func (dict Dictionary) Delete(key string) error {
	_, err := dict.Search(key)

	switch err {
	case ErrNotFound:
		return ErrNotFound
	case nil:
		delete(dict, key)
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
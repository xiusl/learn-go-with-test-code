package _7_map

type Dictionary map[string]string

func Search(dictionary Dictionary, key string) string {
	return dictionary[key]
}

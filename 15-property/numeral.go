package _5_property

import "strings"

func ConvertToRoman(num int) string {
	var res strings.Builder

	for i := 0; i < num; i++ {
		res.WriteString("I")
	}
	return res.String()
}

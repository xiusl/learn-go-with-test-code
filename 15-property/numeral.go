package _5_property

import "strings"

func ConvertToRoman(num int) string {
	var res strings.Builder

	for i := num; i > 0; i-- {
		if i == 5 {
			res.WriteString("V")
			break
		}
		if i == 4 {
			res.WriteString("IV")
			break
		}
		res.WriteString("I")
	}
	return res.String()
}


/*
func ConvertToRoman(num int) string {
	var res strings.Builder

	for i := 0; i < num; i++ {
		res.WriteString("I")
	}
	return res.String()
}
*/
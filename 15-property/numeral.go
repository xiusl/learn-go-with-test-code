package _5_property

import "strings"

type RomanNumeral struct {
	Value int
	Symbol string
}

var allRomanNumerals = []RomanNumeral {
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}


func ConvertToRoman(num int) string {
 	var res strings.Builder

 	for _, numeral := range allRomanNumerals {
 		for num >= numeral.Value {
 			res.WriteString(numeral.Symbol)
 			num -= numeral.Value
		}
	}

	return res.String()
}

func ConvertToArabic(roman string) int {
	total := 0
	for range roman {
		total++
	}
	return total
}
/*
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
*/
/*
func ConvertToRoman(num int) string {
	var res strings.Builder

	for i := 0; i < num; i++ {
		res.WriteString("I")
	}
	return res.String()
}
*/
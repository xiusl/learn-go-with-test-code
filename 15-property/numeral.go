package _5_property

import "strings"

type RomanNumeral struct {
	Value int
	Symbol string
}

type RomanNumerals []RomanNumeral

var allRomanNumerals = RomanNumerals {
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

func (rns RomanNumerals) ValueOf(symbol string) int {
	for _, s := range rns {
		if s.Symbol == symbol {
			return s.Value
		}
	}
	return 0
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
	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		if (i + 1) < len(roman) && symbol == 'I' {
			nextSymbol := roman[i+1]

			potentialNumber := string([]byte{symbol, nextSymbol})

			value := allRomanNumerals.ValueOf(potentialNumber)

			if value != 0 {
				total += value
				i++
			} else {
				total++
			}
		} else {
			total++
		}
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
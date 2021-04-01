package _5_property

import "strings"


// https://github.com/quii/learn-go-with-tests/blob/master/roman-numerals.md

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

func (rns RomanNumerals) ValueOf(symbols ...byte) int {
	symbol := string(symbols)
	for _, s := range rns {
		if s.Symbol == symbol {
			return s.Value
		}
	}
	return 0
}

func (rns RomanNumerals) Exists(symbols ...byte) bool {
	symbol := string(symbols)
	for _, s := range rns {
		if s.Symbol == symbol {
			return true
		}
	}
	return false
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

func ConvertToArabic(roman string) (total int) {
	for _, symbols := range windowedRoman(roman).Symbols() {
		total += allRomanNumerals.ValueOf(symbols...)
	}
	return
}

func couldBeSubtractive(index int, currentSymbol uint8, roman string) bool {
	isSubtractiveSymbol := currentSymbol == 'I' || currentSymbol == 'X' || currentSymbol == 'C'
	return (index + 1) < len(roman) && isSubtractiveSymbol
	//return (index + 1) < len(roman) && currentSymbol == 'I'
}

type windowedRoman string

func (w windowedRoman) Symbols() (symbols [][]byte) {
	for i := 0; i < len(w); i++ {
		symbol := w[i]
		notAtEnd := (i + 1) < len(w)

		if notAtEnd && isSubtractive(symbol) && allRomanNumerals.Exists(symbol, w[i+1]) {
			symbols = append(symbols, []byte{byte(symbol), byte(w[i+1])})
			i++
		} else {
			symbols = append(symbols, []byte{byte(symbol)})
		}
	}
	return
}

func isSubtractive(symbol uint8) bool {
	return symbol == 'I' || symbol == 'X' || symbol == 'C'
}


/*
func ConvertToArabic(roman string) int {
	total := 0
	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		if couldBeSubtractive(i, symbol, roman) {
			if value := allRomanNumerals.ValueOf(symbol, roman[i+1]); value != 0 {
				total += value
				i++
			} else {
				total += allRomanNumerals.ValueOf(symbol)
			}
		} else {
			total += allRomanNumerals.ValueOf(symbol)
		}
	}
	return total
}
*/

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
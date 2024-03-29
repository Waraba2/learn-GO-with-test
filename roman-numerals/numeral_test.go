package main

import (
  "testing"
  "strings"
  "fmt"
  "testing/quick"
  "log"
)

func ConvertToArabic(roman string) (total int) {
	for _, symbols := range windowedRoman(roman).Symbols() {
		total += allRomanNumerals.ValueOf(symbols...)
	}
	return
}

func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

type romanNumeral struct {
	Value  int
	Symbol string
}

type romanNumerals []romanNumeral

func (r romanNumerals) ValueOf(symbols ...byte) int {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}

	return 0
}

func (r romanNumerals) Exists(symbols ...byte) bool {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return true
		}
	}
	return false
}

var allRomanNumerals = romanNumerals{
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

type windowedRoman string

func (w windowedRoman) Symbols() (symbols [][]byte) {
	for i := 0; i < len(w); i++ {
		symbol := w[i]
		notAtEnd := i+1 < len(w)

		if notAtEnd && isSubtractive(symbol) && allRomanNumerals.Exists(symbol, w[i+1]) {
			symbols = append(symbols, []byte{symbol, w[i+1]})
			i++
		} else {
			symbols = append(symbols, []byte{symbol})
		}
	}
	return
}

func isSubtractive(symbol uint8) bool {
	return symbol == 'I' || symbol == 'X' || symbol == 'C'
}

var cases = []struct {
  Arabic int
  Roman  string
}{
  {Arabic: 1, Roman: "I"},
  {Arabic: 2, Roman: "II"},
  {Arabic: 3, Roman: "III"},
  {Arabic: 4, Roman: "IV"},
  {Arabic: 5, Roman: "V"},
  {Arabic: 6, Roman: "VI"},
  {Arabic: 7, Roman: "VII"},
  {Arabic: 8, Roman: "VIII"},
  {Arabic: 9, Roman: "IX"},
  {Arabic: 10, Roman: "X"},
  {Arabic: 14, Roman: "XIV"},
  {Arabic: 18, Roman: "XVIII"},
  {Arabic: 20, Roman: "XX"},
  {Arabic: 39, Roman: "XXXIX"},
  {Arabic: 40, Roman: "XL"},
  {Arabic: 47, Roman: "XLVII"},
  {Arabic: 49, Roman: "XLIX"},
  {Arabic: 50, Roman: "L"},
  {Arabic: 100, Roman: "C"},
  {Arabic: 90, Roman: "XC"},
  {Arabic: 400, Roman: "CD"},
  {Arabic: 500, Roman: "D"},
  {Arabic: 900, Roman: "CM"},
  {Arabic: 1000, Roman: "M"},
  {Arabic: 1984, Roman: "MCMLXXXIV"},
  {Arabic: 3999, Roman: "MMMCMXCIX"},
  {Arabic: 2014, Roman: "MMXIV"},
  {Arabic: 1006, Roman: "MVI"},
  {Arabic: 798, Roman: "DCCXCVIII"},  
}
func TestRomanNumerals(t *testing.T) {

  for _, test := range cases {
    t.Run(fmt.Sprintf("%d gets conveted to %q",test.Arabic, test.Roman), func (t *testing.T) {
      got := ConvertToRoman(test.Arabic)

      if got != test.Roman {
        t.Errorf("got %q, want %q", got, test.Roman)
      }
    })
  }
}

func TestCovertingToArabic(t *testing.T){
  for _, test := range cases[:5] {
    t.Run(fmt.Sprintf("%d gets conveted to %q",test.Arabic, test.Roman), func (t *testing.T) {
      got := ConvertToArabic(test.Roman)

      if got != test.Arabic {
        t.Errorf("got %q, want %q", got, test.Arabic)
      }
    })
  }
}

func TestPropertiesOfConversion(t *testing.T) {
  assertion := func(arabic uint16) bool {
    if arabic > 3999 {
      log.Println(arabic)
      return true
    }
    roman := ConvertToRoman(arabic)
    fromRoman := ConvertToArabic(roman)
    return fromRoman == arabic
  }

	if err := quick.Check(assertion, &quick.Config{
    MaxCount: 1000,
  }); err != nil {
		t.Error("failed checks", err)
	}
}

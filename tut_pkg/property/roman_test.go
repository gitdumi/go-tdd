package property

import (
	"fmt"
	"testing"
	"testing/quick"
)

type testCase struct {
	arabic uint16
	roman  string
}

var cases = []testCase{
	{arabic: 1, roman: "I"},
	{arabic: 2, roman: "II"},
	{arabic: 4, roman: "IV"},
	{arabic: 5, roman: "V"},
	{arabic: 6, roman: "VI"},
	{arabic: 7, roman: "VII"},
	{arabic: 8, roman: "VIII"},
	{arabic: 9, roman: "IX"},
	{arabic: 10, roman: "X"},
	{arabic: 50, roman: "L"},
	{arabic: 100, roman: "C"},
	{arabic: 500, roman: "D"},
	{arabic: 1000, roman: "M"},
	{arabic: 1984, roman: "MCMLXXXIV"},
}

func TestConvertToRoman(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintf("%v gets converted to %v", tt.arabic, tt.roman), func(t *testing.T) {
			got := ConvertToRoman(tt.arabic)
			assertString(t, got, tt.roman)
		})
	}
}

func TestConvertToArabic(t *testing.T) {
	for _, tt := range cases[:5] {
		t.Run(fmt.Sprintf("%v gets converted to %v", tt.roman, tt.arabic), func(t *testing.T) {
			got := ConvertToArabic(tt.roman)
			assertInt(t, got, tt.arabic)
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		t.Log(arabic)
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, &quick.Config{MaxCount: 1000}); err != nil {
		t.Error("failed checks", err)
	}
}

func assertString(t testing.TB, got, want string) {
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertInt(t testing.TB, got, want uint16) {
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

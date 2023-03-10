package roman

import (
	"fmt"
	"testing"
	"testing/quick"
)

var cases = []struct {
	Arabic int
	Roman  string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{7, "VII"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{40, "XL"},
	{47, "XLVII"},
	{49, "XLIX"},
	{50, "L"},
	{100, "C"},
	{90, "XC"},
	{400, "CD"},
	{500, "D"},
	{900, "CM"},
	{1000, "M"},
	{1984, "MCMLXXXIV"},
	{3999, "MMMCMXCIX"},
	{2014, "MMXIV"},
	{1006, "MVI"},
	{798, "DCCXCVIII"},
}

func TestConvertToRoman(t *testing.T) {

	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman),
			func(t *testing.T) {
				got := ConvertToRoman(test.Arabic)

				if got != test.Roman {
					t.Errorf("got %q, want %q", got, test.Roman)
				}
			})

	}
}

func TestConvertToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic),
			func(t *testing.T) {
				got := ConvertToArabic(test.Roman)

				if got != test.Arabic {
					t.Errorf("got %d, want %d", got, test.Arabic)
				}
			})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		t.Log("testing", arabic)
		if arabic > 3999 {
			return true
		}

		roman := ConvertToRoman(int(arabic))
		fromRoman := ConvertToArabic(roman)
		return fromRoman == int(arabic)
	}

	if err := quick.Check(assertion,
		&quick.Config{MaxCount: 1000}); err != nil {
		t.Error("failed checks", err)
	}
}

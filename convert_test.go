package main

import "testing"

// iterateRomanTesting used to checking for some valid roman numerals.
func iterateRomanTesting(t *testing.T, numbers []int, romans []string) {
	for i, num := range numbers {
		if roman := ConvertToRoman(num); roman != romans[i] {
			t.Errorf(
				`ConvertToRoman(%d) = %s, want match for %s`,
				num,
				roman,
				romans[i],
			)
		}
	}
}

// TestOneDigitRoman calls ConvertToRoman with one-digit ordinal numbers.
func TestOneDigitRoman(t *testing.T) {
	var (
		numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		romans  = []string{
			"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX",
		}
	)
	iterateRomanTesting(t, numbers, romans)
}

// TestTwoDigitRoman calls ConvertToRoman with two-digit ordinal numbers.
func TestTwoDigitRoman(t *testing.T) {
	var (
		numbers = []int{10, 14, 15, 19, 20, 40, 50, 90, 99}
		romans  = []string{
			"X", "XIV", "XV", "XIX", "XX", "XL", "L", "XC", "XCIX",
		}
	)
	iterateRomanTesting(t, numbers, romans)
}

// TestUniqueThreeDigitRoman calls ConvertToRoman with three-digit numbers,
// but with unique numbers pattern. e.g. 111, 222, 333, 444, 555, ...
func TestUniqueThreeDigitRoman(t *testing.T) {
	var (
		numbers = []int{111, 222, 333, 444, 555, 666, 777, 888, 999}
		romans  = []string{
			"CXI", "CCXXII", "CCCXXXIII", "CDXLIV",
			"DLV", "DCLXVI", "DCCLXXVII", "DCCCLXXXVIII", "CMXCIX",
		}
	)
	iterateRomanTesting(t, numbers, romans)
}

// TestFourDigitRoman calls ConvertToRoman with four-digit ordinal numbers.
func TestFourDigitRoman(t *testing.T) {
	var (
		numbers = []int{1000, 2000, 3000, 3999}
		romans  = []string{"M", "MM", "MMM", "MMMCMXCIX"}
	)
	iterateRomanTesting(t, numbers, romans)
}

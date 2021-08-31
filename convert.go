package main

import "github.com/study-hary-id/roman-numeral/romannumerals"

// convertToRoman return roman numeral effectively until 3999.
func convertToRoman(num int) (romanNum string) {

	var (
		keyNumerals = [13]int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
		index       = len(keyNumerals) - 1
	)
	for num != 0 {
		maxChance := num / keyNumerals[index]
		num %= keyNumerals[index]

		for maxChance != 0 {
			romanNum += romannumerals.Numerals[keyNumerals[index]]
			maxChance -= 1
		}
		index -= 1
	}
	return romanNum

}

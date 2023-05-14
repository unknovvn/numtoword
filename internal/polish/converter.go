package polish

import (
	"fmt"
	"strings"
)

var scales_biggest_to_smallest = []scale{
	{
		value:                 1_000_000_000_000_000_000,
		singularName:          "trylion",
		pluralSingleDigitName: "tryliony",
		pluralTwoDigitName:    "trylionów",
	},
	{
		value:                 1_000_000_000_000_000,
		singularName:          "biliard",
		pluralSingleDigitName: "biliardy",
		pluralTwoDigitName:    "biliardów",
	},
	{
		value:                 1_000_000_000_000,
		singularName:          "bilion",
		pluralSingleDigitName: "biliony",
		pluralTwoDigitName:    "bilionów",
	},
	{
		value:                 1_000_000_000,
		singularName:          "miliard",
		pluralSingleDigitName: "miliardy",
		pluralTwoDigitName:    "miliardów",
	},
	{
		value:                 1_000_000,
		singularName:          "milion",
		pluralSingleDigitName: "miliony",
		pluralTwoDigitName:    "milionów",
	},
	{
		value:                 1000,
		singularName:          "tysiąc",
		pluralSingleDigitName: "tysiące",
		pluralTwoDigitName:    "tysięcy",
	},
}

var zero_to_nineteen_names = []string{
	"zero",
	"jeden",
	"dwa",
	"trzy",
	"cztery",
	"pięć",
	"sześć",
	"siedem",
	"osiem",
	"dziewięć",
	"dziesięć",
	"jedenaście",
	"dwanaście",
	"trzynaście",
	"czternaście",
	"piętnaście",
	"szesnaście",
	"siedemnaście",
	"osiemnaście",
	"dziewiętnaście",
}

var tens_names = []string{
	"dziesięć",
	"dwadzieścia",
	"trzydzieści",
	"czterdzieści",
	"pięćdziesiąt",
	"sześćdziesiąt",
	"siedemdziesiąt",
	"osiemdziesiąt",
	"dziewięćdziesiąt",
}

var hundreds_names = []string{
	"sto",
	"dwieście",
	"trzysta",
	"czterysta",
	"pięćset",
	"sześćset",
	"siedemset",
	"osiemset",
	"dziewięćset",
}

func NumToWord(number int64) string {
	is_negative := false
	remainder := number

	if number < 0 {
		is_negative = true
		remainder = number * -1
	}

	if remainder < int64(len(zero_to_nineteen_names)) {
		return formatResult(is_negative, zero_to_nineteen_names[remainder])
	}

	var resultArray []string
	for _, scale := range scales_biggest_to_smallest {
		if remainder >= scale.value {
			amount := remainder / scale.value
			resultArray = append(resultArray, scale.format(amount, formatZeroToThousand(amount)))
			remainder -= scale.value * amount
		}
	}

	if remainder > 0 {
		resultArray = append(resultArray, formatZeroToThousand(remainder))
	}

	return formatResult(is_negative, strings.Join(resultArray, " "))
}

func formatResult(negativeNumber bool, number string) string {
	if negativeNumber {
		return "minus " + number
	}

	return number
}

func formatZeroToHundred(number int64) string {
	if number < int64(len(zero_to_nineteen_names)) {
		return zero_to_nineteen_names[number]
	}

	tens := number / 10
	remaining_digits := number % 10

	if remaining_digits > 0 {
		return fmt.Sprintf("%s %s", tens_names[tens-1], zero_to_nineteen_names[remaining_digits])
	}

	return tens_names[tens-1]
}

func formatZeroToThousand(number int64) string {
	if number < int64(len(zero_to_nineteen_names)) {
		return zero_to_nineteen_names[number]
	}

	amount_of_hundreds := number / 100
	remainder := number % 100

	if amount_of_hundreds > 0 {
		if remainder > 0 {
			return fmt.Sprintf(
				"%s %s",
				hundreds_names[amount_of_hundreds-1],
				formatZeroToHundred(remainder))
		}

		return hundreds_names[amount_of_hundreds-1]
	}

	if remainder > 0 {
		return formatZeroToHundred(remainder)
	}

	return ""
}

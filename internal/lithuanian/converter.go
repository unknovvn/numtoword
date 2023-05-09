package lithuanian

import (
	"fmt"
	"strings"
)

var scales_biggest_to_smallest = []scale{
	{
		value:                 1_000_000_000_000_000_000,
		singularName:          "kvintilijonas",
		pluralSingleDigitName: "kvintilijonai",
		pluralTwoDigitName:    "kvintilijonų",
	},
	{
		value:                 1_000_000_000_000_000,
		singularName:          "kvadrilijonas",
		pluralSingleDigitName: "kvadrilijonai",
		pluralTwoDigitName:    "kvadrilijonų",
	},
	{
		value:                 1_000_000_000_000,
		singularName:          "trilijonas",
		pluralSingleDigitName: "trilijonai",
		pluralTwoDigitName:    "trilijonų",
	},
	{
		value:                 1_000_000_000,
		singularName:          "milijardas",
		pluralSingleDigitName: "milijardai",
		pluralTwoDigitName:    "milijardų",
	},
	{
		value:                 1_000_000,
		singularName:          "milijonas",
		pluralSingleDigitName: "milijonai",
		pluralTwoDigitName:    "milijonų",
	},
	{
		value:                 1000,
		singularName:          "tūkstantis",
		pluralSingleDigitName: "tūkstančiai",
		pluralTwoDigitName:    "tūkstančių",
	},
	{
		value:                 100,
		singularName:          "šimtas",
		pluralSingleDigitName: "šimtai",
		pluralTwoDigitName:    "šimtų",
	},
}

var zero_to_nineteen_names = []string{
	"nulis",
	"vienas",
	"du",
	"trys",
	"keturi",
	"penki",
	"šeši",
	"septyni",
	"aštuoni",
	"devyni",
	"dešimt",
	"vienuolika",
	"dvylika",
	"trylika",
	"keturiolika",
	"penkiolika",
	"šešiolika",
	"septyniolika",
	"aštuoniolika",
	"devyniolika",
}

var tens_names = []string{
	"dešimt",
	"dvidešimt",
	"trisdešimt",
	"keturiasdešimt",
	"penkiasdešimt",
	"šešiasdešimt",
	"septyniasdešimt",
	"aštuoniasdešimt",
	"devyniasdešimt",
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

	hundred_scale := scales_biggest_to_smallest[len(scales_biggest_to_smallest)-1]
	remainder := number % 100

	if number > hundred_scale.value {
		amount_of_hundreds := number / hundred_scale.value

		if remainder > 0 {
			return fmt.Sprintf(
				"%s %s",
				hundred_scale.format(amount_of_hundreds, formatZeroToThousand(amount_of_hundreds)),
				formatZeroToHundred(remainder))
		}

		return hundred_scale.format(amount_of_hundreds, formatZeroToThousand(amount_of_hundreds))
	}

	if remainder > 0 {
		return formatZeroToHundred(remainder)
	}

	return ""
}

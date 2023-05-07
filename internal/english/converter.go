package english

import (
	"fmt"
	"strings"
)

var scales_biggest_to_smallest = []Scale{
	{Value: 1_000_000_000_000_000_000, Name: "quintillion"},
	{Value: 1_000_000_000_000_000, Name: "quadrillion"},
	{Value: 1_000_000_000_000, Name: "trillion"},
	{Value: 1_000_000_000, Name: "billion"},
	{Value: 1_000_000, Name: "million"},
	{Value: 1000, Name: "thousand"},
	{Value: 100, Name: "hundred"},
}

var zero_to_nineteen_names = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
	"eleven",
	"twelve",
	"thirteen",
	"fourteen",
	"fifteen",
	"sixteen",
	"seventeen",
	"eighteen",
	"nineteen",
}

var tens_names = []string{
	"ten",
	"twenty",
	"thirty",
	"forty",
	"fifty",
	"sixty",
	"seventy",
	"eighty",
	"ninety",
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
		if remainder >= scale.Value {
			amount := remainder / scale.Value
			resultArray = append(resultArray, fmt.Sprintf("%s %s", formatZeroToThousand(amount), scale.Name))
			remainder -= scale.Value * amount
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
		return tens_names[tens-1] + "-" + zero_to_nineteen_names[remaining_digits]
	}

	return tens_names[tens-1]
}

func formatZeroToThousand(number int64) string {
	if number < int64(len(zero_to_nineteen_names)) {
		return zero_to_nineteen_names[number]
	}

	hundred_scale := scales_biggest_to_smallest[len(scales_biggest_to_smallest)-1]
	remainder := number % 100

	if number > hundred_scale.Value {
		amount_of_hundreds := number / hundred_scale.Value

		if remainder > 0 {
			return fmt.Sprintf("%s %s %s", formatZeroToThousand(amount_of_hundreds), hundred_scale.Name, formatZeroToHundred(remainder))
		}

		return fmt.Sprintf("%s %s", formatZeroToThousand(amount_of_hundreds), hundred_scale.Name)
	}

	if remainder > 0 {
		return formatZeroToHundred(remainder)
	}

	return ""
}

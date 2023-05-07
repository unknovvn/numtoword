package english_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/unknovvn/numtoword/internal/english"
)

func TestNumToWord_CorrectlyConvertsPlusMinusNineteen(t *testing.T) {
	assert.Equal(t, "minus nineteen", english.NumToWord(-19))
	assert.Equal(t, "minus eighteen", english.NumToWord(-18))
	assert.Equal(t, "minus seventeen", english.NumToWord(-17))
	assert.Equal(t, "minus sixteen", english.NumToWord(-16))
	assert.Equal(t, "minus fifteen", english.NumToWord(-15))
	assert.Equal(t, "minus fourteen", english.NumToWord(-14))
	assert.Equal(t, "minus thirteen", english.NumToWord(-13))
	assert.Equal(t, "minus twelve", english.NumToWord(-12))
	assert.Equal(t, "minus eleven", english.NumToWord(-11))
	assert.Equal(t, "minus ten", english.NumToWord(-10))
	assert.Equal(t, "minus nine", english.NumToWord(-9))
	assert.Equal(t, "minus eight", english.NumToWord(-8))
	assert.Equal(t, "minus seven", english.NumToWord(-7))
	assert.Equal(t, "minus six", english.NumToWord(-6))
	assert.Equal(t, "minus five", english.NumToWord(-5))
	assert.Equal(t, "minus four", english.NumToWord(-4))
	assert.Equal(t, "minus three", english.NumToWord(-3))
	assert.Equal(t, "minus two", english.NumToWord(-2))
	assert.Equal(t, "minus one", english.NumToWord(-1))
	assert.Equal(t, "zero", english.NumToWord(0))
	assert.Equal(t, "one", english.NumToWord(1))
	assert.Equal(t, "two", english.NumToWord(2))
	assert.Equal(t, "three", english.NumToWord(3))
	assert.Equal(t, "four", english.NumToWord(4))
	assert.Equal(t, "five", english.NumToWord(5))
	assert.Equal(t, "six", english.NumToWord(6))
	assert.Equal(t, "seven", english.NumToWord(7))
	assert.Equal(t, "eight", english.NumToWord(8))
	assert.Equal(t, "nine", english.NumToWord(9))
	assert.Equal(t, "ten", english.NumToWord(10))
	assert.Equal(t, "eleven", english.NumToWord(11))
	assert.Equal(t, "twelve", english.NumToWord(12))
	assert.Equal(t, "thirteen", english.NumToWord(13))
	assert.Equal(t, "fourteen", english.NumToWord(14))
	assert.Equal(t, "fifteen", english.NumToWord(15))
	assert.Equal(t, "sixteen", english.NumToWord(16))
	assert.Equal(t, "seventeen", english.NumToWord(17))
	assert.Equal(t, "eighteen", english.NumToWord(18))
	assert.Equal(t, "nineteen", english.NumToWord(19))
}

func TestNumToWord_CorrectlyConvertsTens(t *testing.T) {
	assert.Equal(t, "ten", english.NumToWord(10))
	assert.Equal(t, "twenty", english.NumToWord(20))
	assert.Equal(t, "thirty", english.NumToWord(30))
	assert.Equal(t, "forty", english.NumToWord(40))
	assert.Equal(t, "fifty", english.NumToWord(50))
	assert.Equal(t, "sixty", english.NumToWord(60))
	assert.Equal(t, "seventy", english.NumToWord(70))
	assert.Equal(t, "eighty", english.NumToWord(80))
	assert.Equal(t, "ninety", english.NumToWord(90))
}

func TestNumToWord_CorrectlyConvertsDozensWithUnits(t *testing.T) {
	assert.Equal(t, "forty", english.NumToWord(40))
	assert.Equal(t, "forty-one", english.NumToWord(41))
	assert.Equal(t, "forty-two", english.NumToWord(42))
	assert.Equal(t, "forty-three", english.NumToWord(43))
	assert.Equal(t, "forty-four", english.NumToWord(44))
	assert.Equal(t, "forty-five", english.NumToWord(45))
	assert.Equal(t, "forty-six", english.NumToWord(46))
	assert.Equal(t, "forty-seven", english.NumToWord(47))
	assert.Equal(t, "forty-eight", english.NumToWord(48))
	assert.Equal(t, "forty-nine", english.NumToWord(49))
}

func TestNumToWord_CorrectlyConvertsHundreds(t *testing.T) {
	assert.Equal(t, "one hundred", english.NumToWord(100))
	assert.Equal(t, "two hundred", english.NumToWord(200))
	assert.Equal(t, "three hundred", english.NumToWord(300))
	assert.Equal(t, "four hundred", english.NumToWord(400))
	assert.Equal(t, "five hundred", english.NumToWord(500))
	assert.Equal(t, "six hundred", english.NumToWord(600))
	assert.Equal(t, "seven hundred", english.NumToWord(700))
	assert.Equal(t, "eight hundred", english.NumToWord(800))
	assert.Equal(t, "nine hundred", english.NumToWord(900))
}

func TestNumToWord_CorrectlyConvertsThousands(t *testing.T) {
	assert.Equal(t, "one thousand", english.NumToWord(1_000))
	assert.Equal(t, "two thousand", english.NumToWord(2_000))
	assert.Equal(t, "three thousand", english.NumToWord(3_000))
	assert.Equal(t, "four thousand", english.NumToWord(4_000))
	assert.Equal(t, "five thousand", english.NumToWord(5_000))
	assert.Equal(t, "six thousand", english.NumToWord(6_000))
	assert.Equal(t, "seven thousand", english.NumToWord(7_000))
	assert.Equal(t, "eight thousand", english.NumToWord(8_000))
	assert.Equal(t, "nine thousand", english.NumToWord(9_000))
}

func TestNumToWord_CorrectlyConvertsMillions(t *testing.T) {
	assert.Equal(t, "one million", english.NumToWord(1_000_000))
	assert.Equal(t, "two million", english.NumToWord(2_000_000))
	assert.Equal(t, "three million", english.NumToWord(3_000_000))
	assert.Equal(t, "four million", english.NumToWord(4_000_000))
	assert.Equal(t, "five million", english.NumToWord(5_000_000))
	assert.Equal(t, "six million", english.NumToWord(6_000_000))
	assert.Equal(t, "seven million", english.NumToWord(7_000_000))
	assert.Equal(t, "eight million", english.NumToWord(8_000_000))
	assert.Equal(t, "nine million", english.NumToWord(9_000_000))
}

func TestNumToWord_CorrectlyConvertsBillions(t *testing.T) {
	assert.Equal(t, "one billion", english.NumToWord(1_000_000_000))
	assert.Equal(t, "two billion", english.NumToWord(2_000_000_000))
	assert.Equal(t, "three billion", english.NumToWord(3_000_000_000))
	assert.Equal(t, "four billion", english.NumToWord(4_000_000_000))
	assert.Equal(t, "five billion", english.NumToWord(5_000_000_000))
	assert.Equal(t, "six billion", english.NumToWord(6_000_000_000))
	assert.Equal(t, "seven billion", english.NumToWord(7_000_000_000))
	assert.Equal(t, "eight billion", english.NumToWord(8_000_000_000))
	assert.Equal(t, "nine billion", english.NumToWord(9_000_000_000))
}

func TestNumToWord_CorrectlyConvertsTrillions(t *testing.T) {
	assert.Equal(t, "one trillion", english.NumToWord(1_000_000_000_000))
	assert.Equal(t, "two trillion", english.NumToWord(2_000_000_000_000))
	assert.Equal(t, "three trillion", english.NumToWord(3_000_000_000_000))
	assert.Equal(t, "four trillion", english.NumToWord(4_000_000_000_000))
	assert.Equal(t, "five trillion", english.NumToWord(5_000_000_000_000))
	assert.Equal(t, "six trillion", english.NumToWord(6_000_000_000_000))
	assert.Equal(t, "seven trillion", english.NumToWord(7_000_000_000_000))
	assert.Equal(t, "eight trillion", english.NumToWord(8_000_000_000_000))
	assert.Equal(t, "nine trillion", english.NumToWord(9_000_000_000_000))
}

func TestNumToWord_CorrectlyConvertsQuadrillions(t *testing.T) {
	assert.Equal(t, "one quadrillion", english.NumToWord(1_000_000_000_000_000))
	assert.Equal(t, "two quadrillion", english.NumToWord(2_000_000_000_000_000))
	assert.Equal(t, "three quadrillion", english.NumToWord(3_000_000_000_000_000))
	assert.Equal(t, "four quadrillion", english.NumToWord(4_000_000_000_000_000))
	assert.Equal(t, "five quadrillion", english.NumToWord(5_000_000_000_000_000))
	assert.Equal(t, "six quadrillion", english.NumToWord(6_000_000_000_000_000))
	assert.Equal(t, "seven quadrillion", english.NumToWord(7_000_000_000_000_000))
	assert.Equal(t, "eight quadrillion", english.NumToWord(8_000_000_000_000_000))
	assert.Equal(t, "nine quadrillion", english.NumToWord(9_000_000_000_000_000))
}

func TestNumToWord_CorrectlyConvertsQuintillions(t *testing.T) {
	assert.Equal(t, "one quintillion", english.NumToWord(1_000_000_000_000_000_000))
	assert.Equal(t, "two quintillion", english.NumToWord(2_000_000_000_000_000_000))
	assert.Equal(t, "three quintillion", english.NumToWord(3_000_000_000_000_000_000))
	assert.Equal(t, "four quintillion", english.NumToWord(4_000_000_000_000_000_000))
	assert.Equal(t, "five quintillion", english.NumToWord(5_000_000_000_000_000_000))
	assert.Equal(t, "six quintillion", english.NumToWord(6_000_000_000_000_000_000))
	assert.Equal(t, "seven quintillion", english.NumToWord(7_000_000_000_000_000_000))
	assert.Equal(t, "eight quintillion", english.NumToWord(8_000_000_000_000_000_000))
	assert.Equal(t, "nine quintillion", english.NumToWord(9_000_000_000_000_000_000))
}

func TestNumToWord_CorrectlyConvertsCombinedScales(t *testing.T) {
	assert.Equal(t, "four hundred forty-four", english.NumToWord(444))
	assert.Equal(t, "nine hundred ninety-nine", english.NumToWord(999))
	assert.Equal(t, "four thousand four hundred forty-four", english.NumToWord(4_444))
	assert.Equal(t, "nine thousand nine hundred ninety-nine", english.NumToWord(9_999))
	assert.Equal(
		t,
		"four million four hundred forty-four thousand four hundred forty-four",
		english.NumToWord(4_444_444))
	assert.Equal(
		t,
		"nine million nine hundred ninety-nine thousand nine hundred ninety-nine",
		english.NumToWord(9_999_999))
	assert.Equal(
		t,
		"four billion four hundred forty-four million four hundred forty-four thousand four hundred forty-four",
		english.NumToWord(4_444_444_444))
	assert.Equal(
		t,
		"nine billion nine hundred ninety-nine million nine hundred ninety-nine thousand nine hundred ninety-nine",
		english.NumToWord(9_999_999_999))
	assert.Equal(
		t,
		"four trillion four hundred forty-four billion four hundred forty-four million four hundred forty-four thousand four hundred forty-four",
		english.NumToWord(4_444_444_444_444))
	assert.Equal(
		t,
		"nine trillion nine hundred ninety-nine billion nine hundred ninety-nine million nine hundred ninety-nine thousand nine hundred ninety-nine",
		english.NumToWord(9_999_999_999_999))
	assert.Equal(
		t,
		"four quadrillion four hundred forty-four trillion four hundred forty-four billion four hundred forty-four million four hundred forty-four thousand four hundred forty-four",
		english.NumToWord(4_444_444_444_444_444))
	assert.Equal(
		t,
		"nine quadrillion nine hundred ninety-nine trillion nine hundred ninety-nine billion nine hundred ninety-nine million nine hundred ninety-nine thousand nine hundred ninety-nine",
		english.NumToWord(9_999_999_999_999_999))
	assert.Equal(
		t,
		"four quintillion four hundred forty-four quadrillion four hundred forty-four trillion four hundred forty-four billion four hundred forty-four million four hundred forty-four thousand four hundred forty-four",
		english.NumToWord(4_444_444_444_444_444_444))
	assert.Equal(
		t,
		"nine quintillion one hundred ninety-nine quadrillion nine hundred ninety-nine trillion nine hundred ninety-nine billion nine hundred ninety-nine million nine hundred ninety-nine thousand nine hundred ninety-nine",
		english.NumToWord(9_199_999_999_999_999_999))
}

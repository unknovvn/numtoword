package polish_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/unknovvn/numtoword/internal/polish"
)

func TestNumToWord_CorrectlyConvertsPlusMinusNineteen(t *testing.T) {
	assert.Equal(t, "minus dziewiętnaście", polish.NumToWord(-19))
	assert.Equal(t, "minus osiemnaście", polish.NumToWord(-18))
	assert.Equal(t, "minus siedemnaście", polish.NumToWord(-17))
	assert.Equal(t, "minus szesnaście", polish.NumToWord(-16))
	assert.Equal(t, "minus piętnaście", polish.NumToWord(-15))
	assert.Equal(t, "minus czternaście", polish.NumToWord(-14))
	assert.Equal(t, "minus trzynaście", polish.NumToWord(-13))
	assert.Equal(t, "minus dwanaście", polish.NumToWord(-12))
	assert.Equal(t, "minus jedenaście", polish.NumToWord(-11))
	assert.Equal(t, "minus dziesięć", polish.NumToWord(-10))
	assert.Equal(t, "minus dziewięć", polish.NumToWord(-9))
	assert.Equal(t, "minus osiem", polish.NumToWord(-8))
	assert.Equal(t, "minus siedem", polish.NumToWord(-7))
	assert.Equal(t, "minus sześć", polish.NumToWord(-6))
	assert.Equal(t, "minus pięć", polish.NumToWord(-5))
	assert.Equal(t, "minus cztery", polish.NumToWord(-4))
	assert.Equal(t, "minus trzy", polish.NumToWord(-3))
	assert.Equal(t, "minus dwa", polish.NumToWord(-2))
	assert.Equal(t, "minus jeden", polish.NumToWord(-1))
	assert.Equal(t, "zero", polish.NumToWord(0))
	assert.Equal(t, "jeden", polish.NumToWord(1))
	assert.Equal(t, "dwa", polish.NumToWord(2))
	assert.Equal(t, "trzy", polish.NumToWord(3))
	assert.Equal(t, "cztery", polish.NumToWord(4))
	assert.Equal(t, "pięć", polish.NumToWord(5))
	assert.Equal(t, "sześć", polish.NumToWord(6))
	assert.Equal(t, "siedem", polish.NumToWord(7))
	assert.Equal(t, "osiem", polish.NumToWord(8))
	assert.Equal(t, "dziewięć", polish.NumToWord(9))
	assert.Equal(t, "dziesięć", polish.NumToWord(10))
	assert.Equal(t, "jedenaście", polish.NumToWord(11))
	assert.Equal(t, "dwanaście", polish.NumToWord(12))
	assert.Equal(t, "trzynaście", polish.NumToWord(13))
	assert.Equal(t, "czternaście", polish.NumToWord(14))
	assert.Equal(t, "piętnaście", polish.NumToWord(15))
	assert.Equal(t, "szesnaście", polish.NumToWord(16))
	assert.Equal(t, "siedemnaście", polish.NumToWord(17))
	assert.Equal(t, "osiemnaście", polish.NumToWord(18))
	assert.Equal(t, "dziewiętnaście", polish.NumToWord(19))
}

func TestNumToWord_CorrectlyConvertsTens(t *testing.T) {
	assert.Equal(t, "dziesięć", polish.NumToWord(10))
	assert.Equal(t, "dwadzieścia", polish.NumToWord(20))
	assert.Equal(t, "trzydzieści", polish.NumToWord(30))
	assert.Equal(t, "czterdzieści", polish.NumToWord(40))
	assert.Equal(t, "pięćdziesiąt", polish.NumToWord(50))
	assert.Equal(t, "sześćdziesiąt", polish.NumToWord(60))
	assert.Equal(t, "siedemdziesiąt", polish.NumToWord(70))
	assert.Equal(t, "osiemdziesiąt", polish.NumToWord(80))
	assert.Equal(t, "dziewięćdziesiąt", polish.NumToWord(90))
}

func TestNumToWord_CorrectlyConvertsDozensWithUnits(t *testing.T) {
	assert.Equal(t, "czterdzieści", polish.NumToWord(40))
	assert.Equal(t, "czterdzieści jeden", polish.NumToWord(41))
	assert.Equal(t, "czterdzieści dwa", polish.NumToWord(42))
	assert.Equal(t, "czterdzieści trzy", polish.NumToWord(43))
	assert.Equal(t, "czterdzieści cztery", polish.NumToWord(44))
	assert.Equal(t, "czterdzieści pięć", polish.NumToWord(45))
	assert.Equal(t, "czterdzieści sześć", polish.NumToWord(46))
	assert.Equal(t, "czterdzieści siedem", polish.NumToWord(47))
	assert.Equal(t, "czterdzieści osiem", polish.NumToWord(48))
	assert.Equal(t, "czterdzieści dziewięć", polish.NumToWord(49))
}

func TestNumToWord_CorrectlyConvertsHundreds(t *testing.T) {
	assert.Equal(t, "sto", polish.NumToWord(100))
	assert.Equal(t, "dwieście", polish.NumToWord(200))
	assert.Equal(t, "trzysta", polish.NumToWord(300))
	assert.Equal(t, "czterysta", polish.NumToWord(400))
	assert.Equal(t, "pięćset", polish.NumToWord(500))
	assert.Equal(t, "sześćset", polish.NumToWord(600))
	assert.Equal(t, "siedemset", polish.NumToWord(700))
	assert.Equal(t, "osiemset", polish.NumToWord(800))
	assert.Equal(t, "dziewięćset", polish.NumToWord(900))
}

func TestNumToWord_CorrectlyConvertsThousands(t *testing.T) {
	assert.Equal(t, "jeden tysiąc", polish.NumToWord(1_000))
	assert.Equal(t, "dwa tysiące", polish.NumToWord(2_000))
	assert.Equal(t, "trzy tysiące", polish.NumToWord(3_000))
	assert.Equal(t, "cztery tysiące", polish.NumToWord(4_000))
	assert.Equal(t, "pięć tysięcy", polish.NumToWord(5_000))
	assert.Equal(t, "sześć tysięcy", polish.NumToWord(6_000))
	assert.Equal(t, "siedem tysięcy", polish.NumToWord(7_000))
	assert.Equal(t, "osiem tysięcy", polish.NumToWord(8_000))
	assert.Equal(t, "dziewięć tysięcy", polish.NumToWord(9_000))
}

func TestNumToWord_CorrectlyConvertsMillions(t *testing.T) {
	assert.Equal(t, "jeden milion", polish.NumToWord(1_000_000))
	assert.Equal(t, "dwa miliony", polish.NumToWord(2_000_000))
	assert.Equal(t, "trzy miliony", polish.NumToWord(3_000_000))
	assert.Equal(t, "cztery miliony", polish.NumToWord(4_000_000))
	assert.Equal(t, "pięć milionów", polish.NumToWord(5_000_000))
	assert.Equal(t, "sześć milionów", polish.NumToWord(6_000_000))
	assert.Equal(t, "siedem milionów", polish.NumToWord(7_000_000))
	assert.Equal(t, "osiem milionów", polish.NumToWord(8_000_000))
	assert.Equal(t, "dziewięć milionów", polish.NumToWord(9_000_000))
}

func TestNumToWord_CorrectlyConvertsBillions(t *testing.T) {
	assert.Equal(t, "jeden miliard", polish.NumToWord(1_000_000_000))
	assert.Equal(t, "dwa miliardy", polish.NumToWord(2_000_000_000))
	assert.Equal(t, "trzy miliardy", polish.NumToWord(3_000_000_000))
	assert.Equal(t, "cztery miliardy", polish.NumToWord(4_000_000_000))
	assert.Equal(t, "pięć miliardów", polish.NumToWord(5_000_000_000))
	assert.Equal(t, "sześć miliardów", polish.NumToWord(6_000_000_000))
	assert.Equal(t, "siedem miliardów", polish.NumToWord(7_000_000_000))
	assert.Equal(t, "osiem miliardów", polish.NumToWord(8_000_000_000))
	assert.Equal(t, "dziewięć miliardów", polish.NumToWord(9_000_000_000))
}

func TestNumToWord_CorrectlyConvertsTrillions(t *testing.T) {
	assert.Equal(t, "jeden bilion", polish.NumToWord(1_000_000_000_000))
	assert.Equal(t, "dwa biliony", polish.NumToWord(2_000_000_000_000))
	assert.Equal(t, "trzy biliony", polish.NumToWord(3_000_000_000_000))
	assert.Equal(t, "cztery biliony", polish.NumToWord(4_000_000_000_000))
	assert.Equal(t, "pięć bilionów", polish.NumToWord(5_000_000_000_000))
	assert.Equal(t, "sześć bilionów", polish.NumToWord(6_000_000_000_000))
	assert.Equal(t, "siedem bilionów", polish.NumToWord(7_000_000_000_000))
	assert.Equal(t, "osiem bilionów", polish.NumToWord(8_000_000_000_000))
	assert.Equal(t, "dziewięć bilionów", polish.NumToWord(9_000_000_000_000))
}

func TestNumToWord_CorrectlyConvertsQuadrillions(t *testing.T) {
	assert.Equal(t, "jeden biliard", polish.NumToWord(1_000_000_000_000_000))
	assert.Equal(t, "dwa biliardy", polish.NumToWord(2_000_000_000_000_000))
	assert.Equal(t, "trzy biliardy", polish.NumToWord(3_000_000_000_000_000))
	assert.Equal(t, "cztery biliardy", polish.NumToWord(4_000_000_000_000_000))
	assert.Equal(t, "pięć biliardów", polish.NumToWord(5_000_000_000_000_000))
	assert.Equal(t, "sześć biliardów", polish.NumToWord(6_000_000_000_000_000))
	assert.Equal(t, "siedem biliardów", polish.NumToWord(7_000_000_000_000_000))
	assert.Equal(t, "osiem biliardów", polish.NumToWord(8_000_000_000_000_000))
	assert.Equal(t, "dziewięć biliardów", polish.NumToWord(9_000_000_000_000_000))
}

func TestNumToWord_CorrectlyConvertsQuintillions(t *testing.T) {
	assert.Equal(t, "jeden trylion", polish.NumToWord(1_000_000_000_000_000_000))
	assert.Equal(t, "dwa tryliony", polish.NumToWord(2_000_000_000_000_000_000))
	assert.Equal(t, "trzy tryliony", polish.NumToWord(3_000_000_000_000_000_000))
	assert.Equal(t, "cztery tryliony", polish.NumToWord(4_000_000_000_000_000_000))
	assert.Equal(t, "pięć trylionów", polish.NumToWord(5_000_000_000_000_000_000))
	assert.Equal(t, "sześć trylionów", polish.NumToWord(6_000_000_000_000_000_000))
	assert.Equal(t, "siedem trylionów", polish.NumToWord(7_000_000_000_000_000_000))
	assert.Equal(t, "osiem trylionów", polish.NumToWord(8_000_000_000_000_000_000))
	assert.Equal(t, "dziewięć trylionów", polish.NumToWord(9_000_000_000_000_000_000))
}

func TestNumToWord_CorrectlyConvertsCombinedScales(t *testing.T) {
	assert.Equal(t, "czterysta czterdzieści cztery", polish.NumToWord(444))
	assert.Equal(t, "dziewięćset dziewięćdziesiąt dziewięć", polish.NumToWord(999))
	assert.Equal(t, "cztery tysiące czterysta czterdzieści cztery", polish.NumToWord(4_444))
	assert.Equal(t, "dziewięć tysięcy dziewięćset dziewięćdziesiąt dziewięć", polish.NumToWord(9_999))
	assert.Equal(
		t,
		"cztery miliony czterysta czterdzieści cztery tysiące czterysta czterdzieści cztery",
		polish.NumToWord(4_444_444))
	assert.Equal(
		t,
		"dziewięć milionów dziewięćset dziewięćdziesiąt dziewięć tysięcy dziewięćset dziewięćdziesiąt dziewięć",
		polish.NumToWord(9_999_999))
	assert.Equal(
		t,
		"cztery miliardy czterysta czterdzieści cztery miliony czterysta czterdzieści cztery tysiące czterysta czterdzieści cztery",
		polish.NumToWord(4_444_444_444))
	assert.Equal(
		t,
		"dziewięć miliardów dziewięćset dziewięćdziesiąt dziewięć milionów dziewięćset dziewięćdziesiąt dziewięć tysięcy dziewięćset dziewięćdziesiąt dziewięć",
		polish.NumToWord(9_999_999_999))
	assert.Equal(
		t,
		"cztery biliony czterysta czterdzieści cztery miliardy czterysta czterdzieści cztery miliony czterysta czterdzieści cztery tysiące czterysta czterdzieści cztery",
		polish.NumToWord(4_444_444_444_444))
	assert.Equal(
		t,
		"dziewięć bilionów dziewięćset dziewięćdziesiąt dziewięć miliardów dziewięćset dziewięćdziesiąt dziewięć milionów dziewięćset dziewięćdziesiąt dziewięć tysięcy dziewięćset dziewięćdziesiąt dziewięć",
		polish.NumToWord(9_999_999_999_999))
	assert.Equal(
		t,
		"cztery biliardy czterysta czterdzieści cztery biliony czterysta czterdzieści cztery miliardy czterysta czterdzieści cztery miliony czterysta czterdzieści cztery tysiące czterysta czterdzieści cztery",
		polish.NumToWord(4_444_444_444_444_444))
	assert.Equal(
		t,
		"dziewięć biliardów dziewięćset dziewięćdziesiąt dziewięć bilionów dziewięćset dziewięćdziesiąt dziewięć miliardów dziewięćset dziewięćdziesiąt dziewięć milionów dziewięćset dziewięćdziesiąt dziewięć tysięcy dziewięćset dziewięćdziesiąt dziewięć",
		polish.NumToWord(9_999_999_999_999_999))
	assert.Equal(
		t,
		"cztery tryliony czterysta czterdzieści cztery biliardy czterysta czterdzieści cztery biliony czterysta czterdzieści cztery miliardy czterysta czterdzieści cztery miliony czterysta czterdzieści cztery tysiące czterysta czterdzieści cztery",
		polish.NumToWord(4_444_444_444_444_444_444))
	assert.Equal(
		t,
		"dziewięć trylionów sto dziewięćdziesiąt dziewięć biliardów dziewięćset dziewięćdziesiąt dziewięć bilionów dziewięćset dziewięćdziesiąt dziewięć miliardów dziewięćset dziewięćdziesiąt dziewięć milionów dziewięćset dziewięćdziesiąt dziewięć tysięcy dziewięćset dziewięćdziesiąt dziewięć",
		polish.NumToWord(9_199_999_999_999_999_999))
}

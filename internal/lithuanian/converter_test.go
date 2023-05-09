package lithuanian_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/unknovvn/numtoword/internal/lithuanian"
)

func TestNumToWord_CorrectlyConvertsPlusMinusNineteen(t *testing.T) {
	assert.Equal(t, "minus devyniolika", lithuanian.NumToWord(-19))
	assert.Equal(t, "minus aštuoniolika", lithuanian.NumToWord(-18))
	assert.Equal(t, "minus septyniolika", lithuanian.NumToWord(-17))
	assert.Equal(t, "minus šešiolika", lithuanian.NumToWord(-16))
	assert.Equal(t, "minus penkiolika", lithuanian.NumToWord(-15))
	assert.Equal(t, "minus keturiolika", lithuanian.NumToWord(-14))
	assert.Equal(t, "minus trylika", lithuanian.NumToWord(-13))
	assert.Equal(t, "minus dvylika", lithuanian.NumToWord(-12))
	assert.Equal(t, "minus vienuolika", lithuanian.NumToWord(-11))
	assert.Equal(t, "minus dešimt", lithuanian.NumToWord(-10))
	assert.Equal(t, "minus devyni", lithuanian.NumToWord(-9))
	assert.Equal(t, "minus aštuoni", lithuanian.NumToWord(-8))
	assert.Equal(t, "minus septyni", lithuanian.NumToWord(-7))
	assert.Equal(t, "minus šeši", lithuanian.NumToWord(-6))
	assert.Equal(t, "minus penki", lithuanian.NumToWord(-5))
	assert.Equal(t, "minus keturi", lithuanian.NumToWord(-4))
	assert.Equal(t, "minus trys", lithuanian.NumToWord(-3))
	assert.Equal(t, "minus du", lithuanian.NumToWord(-2))
	assert.Equal(t, "minus vienas", lithuanian.NumToWord(-1))
	assert.Equal(t, "nulis", lithuanian.NumToWord(0))
	assert.Equal(t, "vienas", lithuanian.NumToWord(1))
	assert.Equal(t, "du", lithuanian.NumToWord(2))
	assert.Equal(t, "trys", lithuanian.NumToWord(3))
	assert.Equal(t, "keturi", lithuanian.NumToWord(4))
	assert.Equal(t, "penki", lithuanian.NumToWord(5))
	assert.Equal(t, "šeši", lithuanian.NumToWord(6))
	assert.Equal(t, "septyni", lithuanian.NumToWord(7))
	assert.Equal(t, "aštuoni", lithuanian.NumToWord(8))
	assert.Equal(t, "devyni", lithuanian.NumToWord(9))
	assert.Equal(t, "dešimt", lithuanian.NumToWord(10))
	assert.Equal(t, "vienuolika", lithuanian.NumToWord(11))
	assert.Equal(t, "dvylika", lithuanian.NumToWord(12))
	assert.Equal(t, "trylika", lithuanian.NumToWord(13))
	assert.Equal(t, "keturiolika", lithuanian.NumToWord(14))
	assert.Equal(t, "penkiolika", lithuanian.NumToWord(15))
	assert.Equal(t, "šešiolika", lithuanian.NumToWord(16))
	assert.Equal(t, "septyniolika", lithuanian.NumToWord(17))
	assert.Equal(t, "aštuoniolika", lithuanian.NumToWord(18))
	assert.Equal(t, "devyniolika", lithuanian.NumToWord(19))
}

func TestNumToWord_CorrectlyConvertsTens(t *testing.T) {
	assert.Equal(t, "dešimt", lithuanian.NumToWord(10))
	assert.Equal(t, "dvidešimt", lithuanian.NumToWord(20))
	assert.Equal(t, "trisdešimt", lithuanian.NumToWord(30))
	assert.Equal(t, "keturiasdešimt", lithuanian.NumToWord(40))
	assert.Equal(t, "penkiasdešimt", lithuanian.NumToWord(50))
	assert.Equal(t, "šešiasdešimt", lithuanian.NumToWord(60))
	assert.Equal(t, "septyniasdešimt", lithuanian.NumToWord(70))
	assert.Equal(t, "aštuoniasdešimt", lithuanian.NumToWord(80))
	assert.Equal(t, "devyniasdešimt", lithuanian.NumToWord(90))
}

func TestNumToWord_CorrectlyConvertsDozensWithUnits(t *testing.T) {
	assert.Equal(t, "keturiasdešimt", lithuanian.NumToWord(40))
	assert.Equal(t, "keturiasdešimt vienas", lithuanian.NumToWord(41))
	assert.Equal(t, "keturiasdešimt du", lithuanian.NumToWord(42))
	assert.Equal(t, "keturiasdešimt trys", lithuanian.NumToWord(43))
	assert.Equal(t, "keturiasdešimt keturi", lithuanian.NumToWord(44))
	assert.Equal(t, "keturiasdešimt penki", lithuanian.NumToWord(45))
	assert.Equal(t, "keturiasdešimt šeši", lithuanian.NumToWord(46))
	assert.Equal(t, "keturiasdešimt septyni", lithuanian.NumToWord(47))
	assert.Equal(t, "keturiasdešimt aštuoni", lithuanian.NumToWord(48))
	assert.Equal(t, "keturiasdešimt devyni", lithuanian.NumToWord(49))
}

func TestNumToWord_CorrectlyConvertsHundreds(t *testing.T) {
	assert.Equal(t, "vienas šimtas", lithuanian.NumToWord(100))
	assert.Equal(t, "du šimtai", lithuanian.NumToWord(200))
	assert.Equal(t, "trys šimtai", lithuanian.NumToWord(300))
	assert.Equal(t, "keturi šimtai", lithuanian.NumToWord(400))
	assert.Equal(t, "penki šimtai", lithuanian.NumToWord(500))
	assert.Equal(t, "šeši šimtai", lithuanian.NumToWord(600))
	assert.Equal(t, "septyni šimtai", lithuanian.NumToWord(700))
	assert.Equal(t, "aštuoni šimtai", lithuanian.NumToWord(800))
	assert.Equal(t, "devyni šimtai", lithuanian.NumToWord(900))
}

func TestNumToWord_CorrectlyConvertsThousands(t *testing.T) {
	assert.Equal(t, "vienas tūkstantis", lithuanian.NumToWord(1_000))
	assert.Equal(t, "du tūkstančiai", lithuanian.NumToWord(2_000))
	assert.Equal(t, "trys tūkstančiai", lithuanian.NumToWord(3_000))
	assert.Equal(t, "keturi tūkstančiai", lithuanian.NumToWord(4_000))
	assert.Equal(t, "penki tūkstančiai", lithuanian.NumToWord(5_000))
	assert.Equal(t, "šeši tūkstančiai", lithuanian.NumToWord(6_000))
	assert.Equal(t, "septyni tūkstančiai", lithuanian.NumToWord(7_000))
	assert.Equal(t, "aštuoni tūkstančiai", lithuanian.NumToWord(8_000))
	assert.Equal(t, "devyni tūkstančiai", lithuanian.NumToWord(9_000))
}

func TestNumToWord_CorrectlyConvertsMillions(t *testing.T) {
	assert.Equal(t, "vienas milijonas", lithuanian.NumToWord(1_000_000))
	assert.Equal(t, "du milijonai", lithuanian.NumToWord(2_000_000))
	assert.Equal(t, "trys milijonai", lithuanian.NumToWord(3_000_000))
	assert.Equal(t, "keturi milijonai", lithuanian.NumToWord(4_000_000))
	assert.Equal(t, "penki milijonai", lithuanian.NumToWord(5_000_000))
	assert.Equal(t, "šeši milijonai", lithuanian.NumToWord(6_000_000))
	assert.Equal(t, "septyni milijonai", lithuanian.NumToWord(7_000_000))
	assert.Equal(t, "aštuoni milijonai", lithuanian.NumToWord(8_000_000))
	assert.Equal(t, "devyni milijonai", lithuanian.NumToWord(9_000_000))
}

func TestNumToWord_CorrectlyConvertsBillions(t *testing.T) {
	assert.Equal(t, "vienas milijardas", lithuanian.NumToWord(1_000_000_000))
	assert.Equal(t, "du milijardai", lithuanian.NumToWord(2_000_000_000))
	assert.Equal(t, "trys milijardai", lithuanian.NumToWord(3_000_000_000))
	assert.Equal(t, "keturi milijardai", lithuanian.NumToWord(4_000_000_000))
	assert.Equal(t, "penki milijardai", lithuanian.NumToWord(5_000_000_000))
	assert.Equal(t, "šeši milijardai", lithuanian.NumToWord(6_000_000_000))
	assert.Equal(t, "septyni milijardai", lithuanian.NumToWord(7_000_000_000))
	assert.Equal(t, "aštuoni milijardai", lithuanian.NumToWord(8_000_000_000))
	assert.Equal(t, "devyni milijardai", lithuanian.NumToWord(9_000_000_000))
}

func TestNumToWord_CorrectlyConvertsTrillions(t *testing.T) {
	assert.Equal(t, "vienas trilijonas", lithuanian.NumToWord(1_000_000_000_000))
	assert.Equal(t, "du trilijonai", lithuanian.NumToWord(2_000_000_000_000))
	assert.Equal(t, "trys trilijonai", lithuanian.NumToWord(3_000_000_000_000))
	assert.Equal(t, "keturi trilijonai", lithuanian.NumToWord(4_000_000_000_000))
	assert.Equal(t, "penki trilijonai", lithuanian.NumToWord(5_000_000_000_000))
	assert.Equal(t, "šeši trilijonai", lithuanian.NumToWord(6_000_000_000_000))
	assert.Equal(t, "septyni trilijonai", lithuanian.NumToWord(7_000_000_000_000))
	assert.Equal(t, "aštuoni trilijonai", lithuanian.NumToWord(8_000_000_000_000))
	assert.Equal(t, "devyni trilijonai", lithuanian.NumToWord(9_000_000_000_000))
}

func TestNumToWord_CorrectlyConvertsQuadrillions(t *testing.T) {
	assert.Equal(t, "vienas kvadrilijonas", lithuanian.NumToWord(1_000_000_000_000_000))
	assert.Equal(t, "du kvadrilijonai", lithuanian.NumToWord(2_000_000_000_000_000))
	assert.Equal(t, "trys kvadrilijonai", lithuanian.NumToWord(3_000_000_000_000_000))
	assert.Equal(t, "keturi kvadrilijonai", lithuanian.NumToWord(4_000_000_000_000_000))
	assert.Equal(t, "penki kvadrilijonai", lithuanian.NumToWord(5_000_000_000_000_000))
	assert.Equal(t, "šeši kvadrilijonai", lithuanian.NumToWord(6_000_000_000_000_000))
	assert.Equal(t, "septyni kvadrilijonai", lithuanian.NumToWord(7_000_000_000_000_000))
	assert.Equal(t, "aštuoni kvadrilijonai", lithuanian.NumToWord(8_000_000_000_000_000))
	assert.Equal(t, "devyni kvadrilijonai", lithuanian.NumToWord(9_000_000_000_000_000))
}

func TestNumToWord_CorrectlyConvertsQuintillions(t *testing.T) {
	assert.Equal(t, "vienas kvintilijonas", lithuanian.NumToWord(1_000_000_000_000_000_000))
	assert.Equal(t, "du kvintilijonai", lithuanian.NumToWord(2_000_000_000_000_000_000))
	assert.Equal(t, "trys kvintilijonai", lithuanian.NumToWord(3_000_000_000_000_000_000))
	assert.Equal(t, "keturi kvintilijonai", lithuanian.NumToWord(4_000_000_000_000_000_000))
	assert.Equal(t, "penki kvintilijonai", lithuanian.NumToWord(5_000_000_000_000_000_000))
	assert.Equal(t, "šeši kvintilijonai", lithuanian.NumToWord(6_000_000_000_000_000_000))
	assert.Equal(t, "septyni kvintilijonai", lithuanian.NumToWord(7_000_000_000_000_000_000))
	assert.Equal(t, "aštuoni kvintilijonai", lithuanian.NumToWord(8_000_000_000_000_000_000))
	assert.Equal(t, "devyni kvintilijonai", lithuanian.NumToWord(9_000_000_000_000_000_000))
}

func TestNumToWord_CorrectlyConvertsCombinedScales(t *testing.T) {
	assert.Equal(t, "keturi šimtai keturiasdešimt keturi", lithuanian.NumToWord(444))
	assert.Equal(t, "devyni šimtai devyniasdešimt devyni", lithuanian.NumToWord(999))
	assert.Equal(t, "keturi tūkstančiai keturi šimtai keturiasdešimt keturi", lithuanian.NumToWord(4_444))
	assert.Equal(t, "devyni tūkstančiai devyni šimtai devyniasdešimt devyni", lithuanian.NumToWord(9_999))
	assert.Equal(
		t,
		"keturi milijonai keturi šimtai keturiasdešimt keturi tūkstančiai keturi šimtai keturiasdešimt keturi",
		lithuanian.NumToWord(4_444_444))
	assert.Equal(
		t,
		"devyni milijonai devyni šimtai devyniasdešimt devyni tūkstančiai devyni šimtai devyniasdešimt devyni",
		lithuanian.NumToWord(9_999_999))
	assert.Equal(
		t,
		"keturi milijardai keturi šimtai keturiasdešimt keturi milijonai keturi šimtai keturiasdešimt keturi tūkstančiai keturi šimtai keturiasdešimt keturi",
		lithuanian.NumToWord(4_444_444_444))
	assert.Equal(
		t,
		"devyni milijardai devyni šimtai devyniasdešimt devyni milijonai devyni šimtai devyniasdešimt devyni tūkstančiai devyni šimtai devyniasdešimt devyni",
		lithuanian.NumToWord(9_999_999_999))
	assert.Equal(
		t,
		"keturi trilijonai keturi šimtai keturiasdešimt keturi milijardai keturi šimtai keturiasdešimt keturi milijonai keturi šimtai keturiasdešimt keturi tūkstančiai keturi šimtai keturiasdešimt keturi",
		lithuanian.NumToWord(4_444_444_444_444))
	assert.Equal(
		t,
		"devyni trilijonai devyni šimtai devyniasdešimt devyni milijardai devyni šimtai devyniasdešimt devyni milijonai devyni šimtai devyniasdešimt devyni tūkstančiai devyni šimtai devyniasdešimt devyni",
		lithuanian.NumToWord(9_999_999_999_999))
	assert.Equal(
		t,
		"keturi kvadrilijonai keturi šimtai keturiasdešimt keturi trilijonai keturi šimtai keturiasdešimt keturi milijardai keturi šimtai keturiasdešimt keturi milijonai keturi šimtai keturiasdešimt keturi tūkstančiai keturi šimtai keturiasdešimt keturi",
		lithuanian.NumToWord(4_444_444_444_444_444))
	assert.Equal(
		t,
		"devyni kvadrilijonai devyni šimtai devyniasdešimt devyni trilijonai devyni šimtai devyniasdešimt devyni milijardai devyni šimtai devyniasdešimt devyni milijonai devyni šimtai devyniasdešimt devyni tūkstančiai devyni šimtai devyniasdešimt devyni",
		lithuanian.NumToWord(9_999_999_999_999_999))
	assert.Equal(
		t,
		"keturi kvintilijonai keturi šimtai keturiasdešimt keturi kvadrilijonai keturi šimtai keturiasdešimt keturi trilijonai keturi šimtai keturiasdešimt keturi milijardai keturi šimtai keturiasdešimt keturi milijonai keturi šimtai keturiasdešimt keturi tūkstančiai keturi šimtai keturiasdešimt keturi",
		lithuanian.NumToWord(4_444_444_444_444_444_444))
	assert.Equal(
		t,
		"devyni kvintilijonai vienas šimtas devyniasdešimt devyni kvadrilijonai devyni šimtai devyniasdešimt devyni trilijonai devyni šimtai devyniasdešimt devyni milijardai devyni šimtai devyniasdešimt devyni milijonai devyni šimtai devyniasdešimt devyni tūkstančiai devyni šimtai devyniasdešimt devyni",
		lithuanian.NumToWord(9_199_999_999_999_999_999))
}

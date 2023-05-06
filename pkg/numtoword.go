package numtoword

import (
	"github.com/unknovvn/numtoword/internal/english"
	"github.com/unknovvn/numtoword/internal/lithuanian"
	"github.com/unknovvn/numtoword/internal/polish"
	"github.com/unknovvn/numtoword/pkg/language"
)

func Convert(number int64, lang language.Language) string {
	switch lang {
	case language.English:
		return english.NumToWord(number)
	case language.Polish:
		return polish.NumToWord(number)
	case language.Lithuanian:
		return lithuanian.NumToWord(number)
	default:
		return ""
	}
}

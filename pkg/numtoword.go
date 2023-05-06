package numtoword

import (
	"github.com/unknovvn/numtoword/internal/english"
	"github.com/unknovvn/numtoword/internal/lithuanian"
	"github.com/unknovvn/numtoword/internal/polish"
)

type Language uint16

const (
	English Language = iota
	Polish
	Lithuanian
)

func Convert(number int64, language Language) string {
	switch language {
	case English:
		return english.NumToWord(number)
	case Polish:
		return polish.NumToWord(number)
	case Lithuanian:
		return lithuanian.NumToWord(number)
	default:
		return ""
	}
}

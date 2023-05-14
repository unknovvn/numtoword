# numtoword

Package which converts number to text

Languages implemented:

- [x] English
- [x] Lithuanian
- [x] Polish

## Installation

You should already have [Go installed](https://golang.org/doc/install).

```bash

go get -u github.com/unknovvn/numtoword

```

## Usage

```go

import (
	"fmt"

	numtoword "github.com/unknovvn/numtoword/pkg"
	"github.com/unknovvn/numtoword/pkg/language"
)

func main() {
	fmt.Println(numtoword.Convert(1, language.Lithuanian)) // vienas
	fmt.Println(numtoword.Convert(18, language.Polish)) // osiemnaście
	fmt.Println(numtoword.Convert(180, language.English)) // one hundred eighty
	fmt.Println(numtoword.Convert(1_800, language.Lithuanian)) // vienas tūkstantis aštuoni šimtai
	fmt.Println(numtoword.Convert(58_766, language.Polish)) // pięćdziesiąt osiem tysięcy siedemset sześćdziesiąt sześć
	fmt.Println(numtoword.Convert(1_800_004, language.English)) // one million eight hundred thousand four
	fmt.Println(numtoword.Convert(11_800_004_721, language.Lithuanian)) // vienuolika milijardų aštuoni šimtai milijonų keturi tūkstančiai septyni šimtai dvidešimt vienas
	fmt.Println(numtoword.Convert(94_896_131_658_498, language.Polish)) // dziewięćdziesiąt cztery biliony osiemset dziewięćdziesiąt sześć miliardów sto trzydzieści jeden milion sześćset pięćdziesiąt osiem tysięcy czterysta dziewięćdziesiąt osiem
	fmt.Println(numtoword.Convert(8_981_321_654_894_986_543, language.English)) // eight quintillion nine hundred eighty-one quadrillion three hundred twenty-one trillion six hundred fifty-four billion eight hundred ninety-four million nine hundred eighty-six thousand five hundred forty-three
}

```

## Contributors

- [Andžej Korovacki](https://github.com/unknovvn) - creator and maintainer

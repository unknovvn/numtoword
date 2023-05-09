# numtoword

Package which converts number to text

Languages implemented:

- [x] English
- [x] Lithuanian
- [ ] Polish

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
	fmt.Println(numtoword.Convert(1, language.English)) // one

	fmt.Println(numtoword.Convert(18, language.English)) // eighteen

	fmt.Println(numtoword.Convert(180, language.Lithuanian)) // vienas šimtas aštuoniasdešimt

	fmt.Println(numtoword.Convert(1_800, language.Lithuanian)) // vienas tūkstantis aštuoni šimtai

	fmt.Println(numtoword.Convert(1_800_004, language.English)) // one million eight hundred thousand four

	fmt.Println(numtoword.Convert(94_896_131_658_498, language.Lithuanian)) // devyniasdešimt keturi trilijonai aštuoni šimtai devyniasdešimt šeši milijardai vienas šimtas trisdešimt vienas milijonas šeši šimtai penkiasdešimt aštuoni tūkstančiai keturi šimtai devyniasdešimt aštuoni

	fmt.Println(numtoword.Convert(8_981_321_654_894_986_543, language.English)) // eight quintillion nine hundred eighty-one quadrillion three hundred twenty-one trillion six hundred fifty-four billion eight hundred ninety-four million nine hundred eighty-six thousand five hundred forty-three
}

```

## Contributors

- [Andžej Korovacki](https://github.com/unknovvn) - creator and maintainer

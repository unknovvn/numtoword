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
	fmt.Println(numtoword.Convert(1, language.English))
	fmt.Println(numtoword.Convert(18, language.English))
	fmt.Println(numtoword.Convert(180, language.English))
	fmt.Println(numtoword.Convert(1_800, language.English))
	fmt.Println(numtoword.Convert(1_800_004, language.English))
	fmt.Println(numtoword.Convert(94_896_131_658_498, language.English))
	fmt.Println(numtoword.Convert(8_981_321_654_894_986_543, language.English))
}

```

## Contributors

- [Andžej Korovacki](https://github.com/unknovvn) - creator and maintainer

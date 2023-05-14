package polish

import "fmt"

type scale struct {
	value                 int64
	singularName          string
	pluralSingleDigitName string
	pluralTwoDigitName    string
}

func (s scale) format(amount int64, formattedAmount string) string {
	remainingDigits := amount % 10

	if remainingDigits >= 5 {
		return fmt.Sprintf("%s %s", formattedAmount, s.pluralTwoDigitName)
	}

	if remainingDigits == 0 {
		return fmt.Sprintf("%s %s", formattedAmount, s.pluralTwoDigitName)
	}

	if remainingDigits == 1 {
		return fmt.Sprintf("%s %s", formattedAmount, s.singularName)
	}

	return fmt.Sprintf("%s %s", formattedAmount, s.pluralSingleDigitName)
}

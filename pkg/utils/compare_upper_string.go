package utils

import (
	"strings"
)

func CompareUpperString(a string, b string) bool {
	a = strings.ToUpper(a)
	b = strings.ToUpper(b)
	return a == b
}

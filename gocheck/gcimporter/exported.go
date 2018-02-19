package gcimporter

import (
	"unicode"
	"unicode/utf8"
)

func exported(name string) bool {
	ch, _ := utf8.DecodeRuneInString(name)
	return unicode.IsUpper(ch)
}

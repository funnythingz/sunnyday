package sunnyday

import (
	"unicode/utf8"
)

func Truncate(s string, c int, o ...string) string {
	r := []rune(s)
	var p string = "..."

	for _, _o := range o {
		p = _o
	}

	if utf8.RuneCountInString(s) > c {
		return string(r[0:c]) + p
	}
	return s
}

package sunnyday

func Truncate(s string, c int, o ...string) string {
	r := []rune(s)
	var p string = "..."

	for _, _o := range o {
		p = _o
	}

	return string(r[0:c]) + p
}

package sunnyday

import (
	"testing"
)

func TestSunnydayTrunate(t *testing.T) {
	s := "hello, world!"
	st := Truncate(s, 5)
	o := "hello..."

	if st != o {
		t.Errorf("got %v want %v", st, o)
	}
}

func TestSunnydayTrunateOption(t *testing.T) {
	s := "hello, world!"
	st := Truncate(s, 5, "...more")
	o := "hello...more"

	if st != o {
		t.Errorf("got %v want %v", st, o)
	}
}

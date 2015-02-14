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

func TestSunnydayNotTrunate(t *testing.T) {
	s := "hello"
	st := Truncate(s, 10)
	o := "hello"

	if st != o {
		t.Errorf("got %v want %v", st, o)
	}
}

func TestSunnydayTrunateCustomOmission(t *testing.T) {
	s := "hello, world!"
	st := Truncate(s, 5, "...more")
	o := "hello...more"

	if st != o {
		t.Errorf("got %v want %v", st, o)
	}
}

package errutil

import (
	"errors"
	"testing"
)

func TestPrefix(t *testing.T) {
	err := errors.New("error")
	err = Prefix(err, "prefix")
	if str := err.Error(); str != "prefix: error" {
		t.Fatal(str)
	}
}

func TestAppend(t *testing.T) {
	one, two := errors.New("1"), errors.New("2")
	err := Append(one, two)
	if str := err.Error(); str != "1\n2" {
		t.Fatal(str)
	}
}

func TestErrorFunc(t *testing.T) {
	err := Func(func() string { return "fun" })
	if str := err.Error(); str != "fun" {
		t.Fatal(str)
	}
}

// type Printer struct {
// 	prefix string
// }

// func (p Printer) Errorf(format string, a ...any) error {
// 	err := fmt.Errorf(format, a...)
// 	return Prefix(err, p.prefix)
// }

// func TestErrorf(t *testing.T) {
// 	p := Printer{"test"}
// 	var errorf Errorf = p
// 	s := errorf.Errorf("hello %s", "there")
// 	assert.EqualError(t, s, "test: hello there")
// }

type Stringed struct {
	s string
}

func (s Stringed) String() string {
	return s.s
}

func TestJoin(t *testing.T) {
	joined := New("a", Stringed{"b"}, "c")
	if str := joined.Error(); str != "a b c" {
		t.Fatal(str)
	}
}

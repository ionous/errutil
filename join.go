package errutil

import (
	"fmt"
	"strings"
)

var Panic = false

// Fmt maps to fmt.Errorf()
func Fmt(format string, parts ...interface{}) error {
	e := joined(fmt.Sprintf(format, parts...))
	if Panic {
		panic(e)
	}
	return e
}

// New provides the missing fmt.Error()
func New(parts ...interface{}) error {
	// Sprint doesnt add spaces if the parts are strings
	// or, apparently, if parts implement String()
	// how annoying.
	a := make([]string, len(parts))
	for i, el := range parts {
		a[i] = fmt.Sprint(el)
	}
	b := strings.Join(a, " ")
	e := joined(b)
	if Panic {
		panic(e)
	}
	return e
}

type joined string

func (j joined) Error() string {
	return string(j)
}

package errutil

import (
	"errors"
	"fmt"
	"log"
)

// Panic when true triggers on any new errutil error...
// unless it implements the NoPanic interface
var Panic = false

// NoPanic - custom errors can implement this to stop errutil from panicing.
// useful on occasion for provisional error types
type NoPanic interface{ NoPanic() }

// Fmt maps to fmt.Errorf()
func Fmt(format string, parts ...interface{}) error {
	e := fmt.Errorf(format, parts...)
	if panicNow(e) {
		log.Panic(e)
	}
	return e
}

// New provides the missing fmt.Error()
func New(parts ...interface{}) error {
	e := Error(Sprint(parts...))
	if panicNow(e) {
		log.Panic(e)
	}
	return e
}

func panicNow(e error) (okay bool) {
	var i NoPanic
	if Panic && e != nil && !errors.As(e, &i) {
		okay = true
	}
	return
}

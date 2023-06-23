package errutil

import (
	"fmt"
	"log"
)

// Error type which uses a single string as an error.
// Can also be used to create error constants.
// https://dave.cheney.net/2016/04/07/constant-errors
// ex. const MyError = errutil.Error("github.com/a/package/MyError")
type Error string

func (e Error) Error() string { return string(e) }

func (e Error) Is(target error) bool { return target == e }

// Error type which uses a single string as an error
// and which implements the generic NoPanic interface.
type NoPanicError string

func (e NoPanicError) Error() string { return string(e) }

func (e NoPanicError) Is(target error) bool { return target == e }

func (e NoPanicError) NoPanic() {}

// Fmt maps to fmt.Errorf()
func Fmt(format string, parts ...any) error {
	e := fmt.Errorf(format, parts...)
	if panicNow(e) {
		log.Panic(e)
	}
	return e
}

// New captures the passed arguments to build an error.
// Wraps the last element if it's an error.
// The returned error delays evaluating the parts until `Error() string`,
// Relies on the presumption that, in an error state,
// the rest of the program wont be altering pending values anyway.
// If that's not true, then Sprint or Sprintf are better options.
func New(parts ...any) (err error) {
	err = errParts(parts)
	if panicNow(err) {
		log.Panic(err)
	}
	return err
}

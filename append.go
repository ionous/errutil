package errutil

import (
	"errors"
	"log"
	"strings"
)

// Append joins two errors into one; either or both can be nil.
func Append(err error, errs ...error) error {
	for _, e := range errs {
		if e != nil {
			if err == nil {
				err = e
			} else if my, ok := err.(multiError); ok {
				my.errors = append(my.errors, e)
				err = my
			} else {
				err = multiError{[]error{err, e}}
			}
		}
	}
	if panicNow(err) {
		log.Panic(err)
	}
	return err
}

// multiError chains errors together.
// it can expand "recursively", allowing an chain of errors.
// this is similar to https://golang.org/pkg/go/scanner/#ErrorList
type multiError struct {
	errors []error
}

// Error returns the combination of errors separated by a newline.
func (my multiError) Error() (ret string) {
	switch errs := my.errors; len(errs) {
	case 0:
	case 1:
		ret = errs[0].Error()
	default:
		var b strings.Builder
		str := errs[0].Error()
		b.WriteString(str)
		for _, e := range errs[1:] {
			b.WriteString("\n")
			str := e.Error()
			b.WriteString(str)
		}
		ret = b.String()
	}
	return
}

func (my multiError) As(i interface{}) (ret bool) {
	for _, e := range my.errors {
		if errors.As(e, i) {
			ret = true
			break
		}
	}
	return
}

func (my multiError) Is(e error) (ret bool) {
	for i, cnt := 0, len(my.errors); i < cnt; i++ {
		if rev := my.errors[cnt-1-i]; errors.Is(rev, e) {
			ret = true
			break
		}
	}
	return
}

// PrintErrors is a utility function that prints a list of errors to w,
// one error per line, if the err parameter is a MultiError. Otherwise
// it prints the err string.
// w is not a writer b/c log.Writer() doesnt generate log headers automatically
// and note log.Println and fmt.Println have different signatures, so yay.
// we use the simplest form possible here.
func PrintErrors(err error, w func(s string)) {
	if list, ok := err.(multiError); ok {
		for _, e := range list.errors {
			w(e.Error())
		}
	} else if err != nil {
		w(err.Error())
	}
}

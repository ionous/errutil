package errutil

import "strings"

// Append joins two errors into one; either or both can be nil.
func Append(err error, errs ...error) error {
	for _, e := range errs {
		if e != nil {
			if err == nil {
				err = e
			} else if my, ok := err.(multiError); ok {
				my.errors = append(my.errors, e)
			} else {
				err = multiError{[]error{err, e}}
			}
		}
	}
	return err
}

// multiError chains errors together.
// it can expand "recursively", allowing an chain of errors.
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

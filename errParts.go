package errutil

import (
	"errors"
)

// helper for errutil.New
type errParts []any

func (ep errParts) Error() string {
	return Sprint(ep...)
}

func (ep errParts) As(i any) bool {
	return errors.As(ep.last(), i)
}

func (ep errParts) Is(e error) bool {
	return errors.Is(ep.last(), e)
}

func (ep errParts) last() (err error) {
	if cnt := len(ep); cnt > 0 {
		err, _ = ep[cnt-1].(error)
	}
	return
}

package errutil

import "errors"

// Panic when true triggers on any new errutil error...
// unless it implements the NoPanic interface
var Panic = false

// NoPanic - custom errors can implement this to stop errutil from panicing.
// useful on occasion for provisional error types
type NoPanic interface{ NoPanic() }

func panicNow(e error) (okay bool) {
	var i NoPanic
	if Panic && e != nil && !errors.As(e, &i) {
		okay = true
	}
	return
}

package errutil

import (
	"fmt"
	"strings"
)

// may change this eventually to use a lighterweight fmt
// maybe using pkg/internal/reflectlite for %T handling.
var Sprintf = fmt.Sprintf

func Sprint(parts ...interface{}) string {
	// Sprint doesnt add spaces if the parts are strings
	// or, apparently, if parts implement String()
	// how annoying.
	a := make([]string, len(parts))
	for i, el := range parts {
		a[i] = fmt.Sprint(el)
	}
	return strings.Join(a, " ")
}

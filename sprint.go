package errutil

import (
	"fmt"
	"strings"
)

// Sprintf acts like its package fmt counterpart.
// It removes a direct dependency on fmt to make usages other uses of fmt easier to find.
// It may eventually change to use a lighterweight formatting,
// ex. maybe using pkg/internal/reflectlite ( for %T handling. )
var Sprintf = fmt.Sprintf

// Fprintf acts like its package fmt counterpart.
// See errutil.Sprintf for notses.
var Fprintf = fmt.Fprintf

// Sprint fixes the fact that package fmt Sprint doesnt add spaces
// if the parts themselves are strings or if the parts implement Stringer.
// When printing errors, who knows what the types implement;
// we want the error formatting to look good regardless.
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

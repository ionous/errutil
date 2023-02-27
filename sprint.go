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
func Sprint(parts ...any) string {
	var b strings.Builder
	for i, el := range parts {
		if i > 0 {
			b.WriteRune(' ')
		}
		b.WriteString(fmt.Sprint(el))
	}
	return b.String()
}

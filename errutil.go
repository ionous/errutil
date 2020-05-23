package errutil

// Panic when true triggers on any new errutil error.
var Panic = false

// Fmt maps to fmt.Errorf()
func Fmt(format string, parts ...interface{}) error {
	e := Error(Sprintf(format, parts...))
	if Panic {
		panic(e)
	}
	return e
}

// New provides the missing fmt.Error()
func New(parts ...interface{}) error {
	e := Error(Sprint(parts...))
	if Panic {
		panic(e)
	}
	return e
}

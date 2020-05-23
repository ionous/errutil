package errutil

// Error type which uses a single string as an error.
// Can also be used to create error constants.
// https://dave.cheney.net/2016/04/07/constant-errors
// ex. const MyError = errutil.Error("github.com/a/package/MyError")
type Error string

func (e Error) Error() string { return string(e) }

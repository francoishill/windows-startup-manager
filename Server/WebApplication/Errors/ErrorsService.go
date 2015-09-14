package Errors

type ErrorService interface {
	PanicClientErrorLocal(e interface{})
	PanicClientErrorLocal_FormattedString(format string, args ...interface{})
}

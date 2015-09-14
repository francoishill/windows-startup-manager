package DefaultErrorService

import (
	"fmt"
	. "github.com/francoishill/windows-startup-manager/Server/WebApplication/Errors"
)

type defaultErrorService struct{}

type tmpClientError struct {
	Error interface{}
}

func (d *defaultErrorService) PanicClientErrorLocal(e interface{}) {
	panic(&tmpClientError{e})
}

func (d *defaultErrorService) PanicClientErrorLocal_FormattedString(format string, args ...interface{}) {
	d.PanicClientErrorLocal(fmt.Errorf(format, args...))
}

func New() ErrorService {
	return &defaultErrorService{}
}

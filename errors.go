package errors

import "fmt"

type Params map[string]interface{}

const (
	ErrParseTmplError = -1
	ErrExecTmpleError = -2
)

var (
	errorTemplate  = make(map[string]*ErrTemplate)
	errCodeDefined = make(map[string]bool)
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	ID      string `json:"id"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%d %s", e.Code, e.Message)
}

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
	Message string `json:"msg"`
	ID      string `json:"id,omitempty"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%d %s", e.Code, e.Message)
}

func NewError(code int, message string) Error {
	key := fmt.Sprintf("%d", code)
	if _, exist := errCodeDefined[key]; exist {
		panic(fmt.Sprintf("error code %s already exist", key))
	}

	errCodeDefined[key] = true

	return Error{Code: code, Message: message}
}

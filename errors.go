package errors

import "fmt"

type Params map[string]interface{}

const (
	ErrParseTmplError = -1
	ErrExecTmpleError = -2
)

var (
	errorTemplate  = make(map[string]*ErrTemplate)
	errCodeDefined = make(map[string]interface{})
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	ID      string `json:"id,omitempty"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d %s", e.Code, e.Message)
}

func NewError(code int, message string) *Error {
	key := fmt.Sprintf("%d", code)
	if _, exist := errCodeDefined[key]; exist {
		panic(fmt.Sprintf("error code %s already exist", key))
	}

	err := &Error{Code: code, Message: message}
	errCodeDefined[key] = err

	return err
}

func GetError(code int) *Error {
	if err, exist := errCodeDefined[fmt.Sprintf("%d", code)]; exist {
		if e, ok := err.(*Error); ok {
			return e
		}
	}

	return nil
}

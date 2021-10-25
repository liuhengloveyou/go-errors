package errors

import (
	"bytes"
	"fmt"
	"hash/crc32"
	"text/template"
	"time"
)

type ErrTemplate struct {
	code     int
	template string
}

func TN(code int, template string) *ErrTemplate {
	key := fmt.Sprintf("%d", code)
	if _, exist := errCodeDefined[key]; exist {
		strErr := fmt.Sprintf("error code %s already exist", key)
		panic(strErr)
	}

	errT := &ErrTemplate{code: code, template: template}
	errCodeDefined[key] = errT

	return errT
}

func (p *ErrTemplate) New(v ...Params) (err error) {
	params := Params{}

	for _, param := range v {
		for pn, pv := range param {
			params[pn] = pv
		}
	}

	strCRCErrId := fmt.Sprintf("%0X", crc32.ChecksumIEEE([]byte(fmt.Sprintf("%d.%s.%d", p.code, p.template, time.Now().UnixNano()))))
	if len(strCRCErrId) > 7 {
		strCRCErrId = strCRCErrId[0:7]
	}

	var tpl *ErrTemplate = p

	t, e := template.New(fmt.Sprintf("%d", p.code)).Parse(tpl.template)
	if e != nil {
		return &Error{
			ID:      strCRCErrId,
			Code:    ErrParseTmplError,
			Message: fmt.Sprintf("parser error template failed, code: %d, error: %s", tpl.code, e.Error()),
		}
	}

	var buf bytes.Buffer
	if e := t.Execute(&buf, params); e != nil {
		return &Error{
			ID:      strCRCErrId,
			Code:    ErrExecTmpleError,
			Message: fmt.Sprintf("execute template failed, code: %d, error: %s", tpl.code, e.Error()),
		}
	}

	return &Error{
		ID:      strCRCErrId,
		Code:    tpl.code,
		Message: buf.String(),
	}
}

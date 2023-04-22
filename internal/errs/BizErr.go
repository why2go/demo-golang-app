package errs

import "fmt"

type BizErr struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func NewBizErr(code, msg string) *BizErr {
	return &BizErr{
		Code: code,
		Msg:  msg,
	}
}

func (e BizErr) Error() string {
	return fmt.Sprintf("[%s], %s", e.Code, e.Msg)
}

var (
	ErrMalformedParams = NewBizErr("e1001", "malformed params")
)

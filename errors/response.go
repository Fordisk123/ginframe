package errors

import (
	"fmt"
)

type ResponseError interface {
	error
}

type RequestError struct {
	RtnCode     string `json:"rtnCode"`
	RtnMsg      string `json:"rtnMsg"`
	DetailError string `json:"detailError"`
}

func (r RequestError) Error() string {

	switch r.RtnCode {
	case "000400":
		{
			return fmt.Sprintf("BadRequest! 错误码:'%s',错误原因：%s, 详细错误:%s", r.RtnCode, r.RtnMsg, r.DetailError)
		}
	case "000500":
		{
			return fmt.Sprintf("Internal Server Error! 错误码:'%s',错误原因：%s, 详细错误:%s", r.RtnCode, r.RtnMsg, r.DetailError)
		}

	default:
		return fmt.Sprintf("Unkown Error! 错误码:'%s',错误原因：%s, 详细错误:%s", r.RtnCode, r.RtnMsg, r.DetailError)
	}

}

func NewBadRequestError(msg string, err error) *RequestError {
	return &RequestError{
		RtnCode:     "000400",
		RtnMsg:      msg,
		DetailError: err.Error(),
	}
}

func NewInternalServerError(msg string, err error) *RequestError {
	return &RequestError{
		RtnCode:     "000500",
		RtnMsg:      msg,
		DetailError: err.Error(),
	}
}

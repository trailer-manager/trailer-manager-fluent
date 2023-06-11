package common

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
)

type Error struct {
	Id      string `json:"errorId"`   // (Mandatory) 에러 ID
	Code    string `json:"errorCode"` // (Mandatory) 에러 코드
	Message string `json:"message"`   // (Mandatory) 메시지
}

func (err *Error) Error() string {
	return fmt.Sprintf("[%s] %s", err.Code, err.Message)
}

// 내부 에러
var (
	ConfigErr = fmt.Errorf("Invalid Configuration Error")
)

// 웹 에러
const (
	ServerErrSuccess          = 2000
	ServerErrBadRequest       = 4000
	ServerErrInvalidParameter = 4001
	ServerErrInvalidValues    = 4002
	ServerErrUnauthorized     = 4010
	ServerErrForbidden        = 4030
	ServerErrNotFound         = 4040
	ServerErrRequestTimeout   = 4080
	ServerErrInternalServer   = 5000
)

var (
	WebErrors = map[int]Error{
		ServerErrBadRequest: {
			Code:    fmt.Sprintf("%d", ServerErrBadRequest),
			Message: "Bad Request",
		}, ServerErrInvalidParameter: {
			Code:    fmt.Sprintf("%d", ServerErrInvalidParameter),
			Message: "Invalid Parameter",
		}, ServerErrInvalidValues: {
			Code:    fmt.Sprintf("%d", ServerErrInvalidValues),
			Message: "Invalid Values",
		},
		ServerErrUnauthorized: {
			Code:    fmt.Sprintf("%d", ServerErrUnauthorized),
			Message: "Unauthorized Request",
		},
		ServerErrForbidden: {
			Code:    fmt.Sprintf("%d", ServerErrForbidden),
			Message: "Forbidden",
		},
		ServerErrNotFound: {
			Code:    fmt.Sprintf("%d", ServerErrNotFound),
			Message: "Not Found",
		},
		ServerErrRequestTimeout: {
			Code:    fmt.Sprintf("%d", ServerErrRequestTimeout),
			Message: "Request Timeout",
		},
		ServerErrInternalServer: {
			Code:    fmt.Sprintf("%d", ServerErrInternalServer),
			Message: "Internal Server Error",
		},
	}
)

func NewServerError(c interface{}, errCode int) (err *Error) {
	ctxType, ctxGo, ctxEcho := getContextType(c)

	err = new(Error)
	err.Code = WebErrors[errCode].Code
	err.Message = WebErrors[errCode].Message

	if ctxType == ContextTypeGo && ctxGo != nil {
		traceId, _ := ctxGo.Value(HeaderTransactionId).(string)
		err.Id = traceId
	} else if ctxType == ContextTypeEcho && ctxEcho != nil {
		err.Id = ctxEcho.Request().Header.Get(HeaderTransactionId)
	}

	return
}

func NewServerErrorWithMessage(c interface{}, errCode int, message string) (err *Error) {
	ctxType, ctxGo, ctxEcho := getContextType(c)

	err = new(Error)
	err.Code = WebErrors[errCode].Code
	err.Message = WebErrors[errCode].Message
	if message != "" {
		err.Message = fmt.Sprintf("%s - %s", err.Message, message)
	}

	if ctxType == ContextTypeGo && ctxGo != nil {
		traceId, _ := ctxGo.Value(HeaderTransactionId).(string)
		err.Id = traceId
	} else if ctxType == ContextTypeEcho && ctxEcho != nil {
		err.Id = ctxEcho.Request().Header.Get(HeaderTransactionId)
	}

	return
}

func (err *Error) IsError() bool {
	return err != nil
}

func getContextType(c interface{}) (ctxType int, ctxGo context.Context, ctxEcho echo.Context) {
	ctxType = ContextTypeNone
	if c != nil {
		ok := false
		if ctxGo, ok = c.(context.Context); ok {
			ctxType = ContextTypeGo
		} else if ctxEcho, ok = c.(echo.Context); ok {
			ctxType = ContextTypeEcho
		}
	}
	return
}

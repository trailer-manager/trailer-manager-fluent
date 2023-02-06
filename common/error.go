package common

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
)

type TMError struct {
	Id      string `json:"errorId"`   // (Mandatory) 에러 ID
	Code    string `json:"errorCode"` // (Mandatory) 에러 코드
	Message string `json:"message"`   // (Mandatory) 메시지
}

func (tmErr *TMError) Error() string {
	return fmt.Sprintf("[%s] %s", tmErr.Code, tmErr.Message)
}

// 내부 에러
var (
	ConfigErr = fmt.Errorf("Invalid Configuration Error")
)

// 웹 에러
const (
	WebErrSuccess        = 2000
	WebErrBadRequest     = 4000
	WebErrUnauthorized   = 4010
	WebErrForbidden      = 4030
	WebErrNotFound       = 4040
	WebErrInternalServer = 5000
)

var (
	WebErrors = map[int]TMError{
		WebErrBadRequest: {
			Code:    fmt.Sprintf("%d", WebErrBadRequest),
			Message: "Bad Request",
		},
		WebErrUnauthorized: {
			Code:    fmt.Sprintf("%d", WebErrUnauthorized),
			Message: "Unauthorized Request",
		},
		WebErrForbidden: {
			Code:    fmt.Sprintf("%d", WebErrForbidden),
			Message: "Forbidden",
		},
		WebErrNotFound: {
			Code:    fmt.Sprintf("%d", WebErrNotFound),
			Message: "Not Found",
		},
		WebErrInternalServer: {
			Code:    fmt.Sprintf("%d", WebErrInternalServer),
			Message: "Internal Server Error",
		},
	}
)

func GetWebError(c interface{}, errCode int) (err *TMError) {
	ctxType, ctxGo, ctxEcho := getContextType(c)

	err = new(TMError)
	err.Code = WebErrors[errCode].Code
	err.Message = WebErrors[errCode].Message

	if ctxType == ContextTypeGo && ctxGo != nil {
		traceId, _ := ctxGo.Value("traceId").(string)
		err.Id = traceId
	} else if ctxType == ContextTypeEcho && ctxEcho != nil {
		err.Id = ctxEcho.Request().Header.Get("traceId")
	}

	return
}

func GetWebErrorWithMessage(c interface{}, errCode int, message string) (err *TMError) {
	ctxType, ctxGo, ctxEcho := getContextType(c)

	err = new(TMError)
	err.Code = WebErrors[errCode].Code
	err.Message = WebErrors[errCode].Message
	if message != "" {
		err.Message = fmt.Sprintf("%s - %s", err.Message, message)
	}

	if ctxType == ContextTypeGo && ctxGo != nil {
		traceId, _ := ctxGo.Value("traceId").(string)
		err.Id = traceId
	} else if ctxType == ContextTypeEcho && ctxEcho != nil {
		err.Id = ctxEcho.Request().Header.Get("traceId")
	}

	return
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
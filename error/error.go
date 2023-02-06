package error

import "fmt"

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
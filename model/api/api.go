package model

type Response struct {
	Result Result `json:"result"`
}

type Result struct {
	StatusCode    int    `json:"statusCode"`
	ErrorCode     string `json:"errCode,omitempty"`
	TransactionId string `json:"tid"`
	Message       string `json:"msg,omitempty"`
}

package model

type Response struct {
	Result Result `json:"result"`
}

type Result struct {
	StatusCode    int    `json:"statusCode"`
	TransactionId string `json:"tid"`
}

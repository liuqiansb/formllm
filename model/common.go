package model

type Result struct {
	Code int64       `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

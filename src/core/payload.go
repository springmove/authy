package core

type RequestError struct {
	Code string `json:"data"`
	Msg  string `json:"msg"`
}

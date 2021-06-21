package models

type RequestError struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}

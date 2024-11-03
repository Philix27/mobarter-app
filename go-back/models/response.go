package models

type Response struct {
	Code    int `json:"code"`
	Status  int `json:"status"`
	Message int `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

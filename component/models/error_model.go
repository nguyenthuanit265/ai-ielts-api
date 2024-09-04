package models

type AIIeltsError struct {
	IsError bool   `json:"-,omitempty" default:"false"`
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Error   error  `json:"error,omitempty"`
}

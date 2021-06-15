package model

type ApiError struct {
	StatusCode int
	ErrorType  ErrorType
}

type ApiSuccess struct {
	StatusCode int
	Data       interface{}
}

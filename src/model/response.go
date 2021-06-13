package model

type ApiError struct {
	StatusCode int
	ErrorType  string
}

type ApiSuccess struct {
	StatusCode int
	Data       interface{}
}

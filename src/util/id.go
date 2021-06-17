package util

import "github.com/google/uuid"

func GetNewTodoId() uuid.UUID {
	return uuid.New()
}

func GetNewUserId() uuid.UUID {
	return uuid.New()
}

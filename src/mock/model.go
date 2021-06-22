package mock

import (
	"fmt"

	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/google/uuid"
)

func GetMockUser() model.AddUser {
	str := uuid.New().String()
	return model.AddUser{
		Name:  str,
		Email: fmt.Sprintf(`%s@test.com`, str),
		SignIn: model.SignIn{
			Account:  str,
			Password: str,
		},
	}
}

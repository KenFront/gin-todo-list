package mock

import (
	"testing"

	"github.com/KenFront/gin-todo-list/src/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var gormDB *gorm.DB

func newGorm() *gorm.DB {
	gormDB, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	if err := gormDB.AutoMigrate(&model.Todo{}); err != nil {
		panic(err)
	}

	if err := gormDB.AutoMigrate(&model.User{}); err != nil {
		panic(err)
	}
	return gormDB
}

func GetMockGorm(t *testing.T) *gorm.DB {
	if gormDB == nil {
		gormDB = newGorm()
	}
	return gormDB
}

package mock

import (
	"testing"

	"github.com/KenFront/gin-todo-list/src/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetMockGorm(t *testing.T) *gorm.DB {
	gormDB, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

	if err != nil {
		t.Fatalf("gorm error: %s", err)
	}
	gormDB.AutoMigrate(&model.Todo{})
	gormDB.AutoMigrate(&model.User{})

	return gormDB
}

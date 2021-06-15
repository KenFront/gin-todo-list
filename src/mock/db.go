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

	if err := gormDB.AutoMigrate(&model.Todo{}); err != nil {
		t.Fatalf("migreate todo error: %s", err)
	}

	if err := gormDB.AutoMigrate(&model.User{}); err != nil {
		t.Fatalf("migreate user error: %s", err)
	}

	return gormDB
}

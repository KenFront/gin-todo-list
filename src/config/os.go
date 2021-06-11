package config

import (
	"os"
)

func InitOs() {
	os.Setenv("TZ", "0")
}

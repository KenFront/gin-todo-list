package util

import "time"

func GetAuthDuration() int {
	const hour = 60 * 60

	return 1 * hour
}

func GetAuthExpiresAt() int64 {
	return time.Now().Add(1 * time.Hour).UTC().Unix()
}

func GetAuthNow() int64 {
	return time.Now().UTC().Unix()
}

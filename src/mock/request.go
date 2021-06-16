package mock

import (
	"bytes"
	"encoding/json"
	"io"
)

func GetRequsetBody(payload interface{}) io.ReadCloser {
	jsonbytes, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	return io.NopCloser(bytes.NewBuffer(jsonbytes))
}

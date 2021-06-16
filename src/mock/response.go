package mock

import (
	"encoding/json"
	"net/http/httptest"
)

func GetResponse() *httptest.ResponseRecorder {
	return httptest.NewRecorder()
}

func GetResponseBody(body []byte, data interface{}) {
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}
}

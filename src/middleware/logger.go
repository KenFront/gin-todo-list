package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"time"

	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type logBase struct {
	Ip            string      `json:"ip"`
	UserId        uuid.UUID   `json:"user_id"`
	StartAt       time.Time   `json:"start_at"`
	EndAt         time.Time   `json:"end_at"`
	StatusCode    int         `json:"status_code"`
	Method        string      `json:"mathod"`
	Path          string      `json:"path"`
	Handlers      []string    `json:"handlers"`
	ErrorMessages []string    `json:"error_message"`
	Payload       interface{} `json:"payload"`
}

func getPath(c *gin.Context) string {
	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery
	if raw == "" {
		return path
	}

	return path + "?" + raw
}

func getPayload(c *gin.Context) string {
	body := c.Request.Body
	data, _ := ioutil.ReadAll(body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
	return string(data)
}

func getPrettyLog(log logBase) string {
	formated, err := json.MarshalIndent(log, "", "  ")
	if err != nil {
		fmt.Println(log)
	}
	result := string(formated)
	return result
}

func hideSecurityPayload(val string) interface{} {
	securities := `\"(password|account)\": \".*\"`
	re := regexp.MustCompile(securities)
	result := re.ReplaceAllString(val, `"$1": "******"`)

	var data interface{}
	if err := json.Unmarshal([]byte(result), &data); err != nil {
		fmt.Println(err)
	}
	return data
}

func customLogger(c *gin.Context) {
	startAt := time.Now()
	path := getPath(c)
	payload := hideSecurityPayload(getPayload(c))

	c.Next()

	endAt := time.Now()
	userId, _ := util.GetUserIdByToken(c)
	errorMessages := c.Errors.Errors()

	log := logBase{
		Ip:            c.ClientIP(),
		UserId:        userId,
		StartAt:       startAt,
		EndAt:         endAt,
		StatusCode:    c.Writer.Status(),
		Method:        c.Request.Method,
		Path:          path,
		Handlers:      c.HandlerNames(),
		ErrorMessages: c.Errors.Errors(),
		Payload:       payload,
	}

	prettyLog := getPrettyLog(log)

	if len(errorMessages) == 0 {
		fmt.Println(color.HiCyanString(prettyLog))
	} else {
		fmt.Println(color.HiRedString(prettyLog))
	}
}

func UseCustomLogger(r *gin.Engine) {
	r.Use(customLogger)
}

func UseLogger(r *gin.Engine) {
	r.Use(gin.Logger())
}

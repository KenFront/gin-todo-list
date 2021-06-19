package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"time"

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
	Query         string      `json:"query"`
	Handlers      []string    `json:"handlers"`
	ErrorMessages []string    `json:"error_message"`
	Payload       interface{} `json:"payload"`
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
	path := c.Request.URL.Path
	query := c.Request.URL.RawQuery
	payload := hideSecurityPayload(getPayload(c))

	c.Next()

	endAt := time.Now()

	userIdByContext, isExist := c.Get("userId")
	var userId uuid.UUID
	if isExist {
		userId = userIdByContext.(uuid.UUID)
	} else {
		userId = uuid.Nil
	}

	errorMessages := c.Errors.Errors()

	go func() {
		log := logBase{
			Ip:            c.ClientIP(),
			UserId:        userId,
			StartAt:       startAt,
			EndAt:         endAt,
			StatusCode:    c.Writer.Status(),
			Method:        c.Request.Method,
			Path:          path,
			Query:         query,
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
	}()
}

func UseCustomLogger(r *gin.Engine) {
	r.Use(customLogger)
}

func UseLogger(r *gin.Engine) {
	r.Use(gin.Logger())
}

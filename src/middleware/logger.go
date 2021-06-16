package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type logBase struct {
	Ip           string      `json:"ip"`
	UserId       uuid.UUID   `json:"user_id"`
	StartAt      time.Time   `json:"start_at"`
	EndAt        time.Time   `json:"end_at"`
	StatusCode   int         `json:"status_code"`
	Method       string      `json:"mathod"`
	Path         string      `json:"path"`
	Handlers     []string    `json:"handlers"`
	ErrorMessage []string    `json:"error_message"`
	Payload      interface{} `json:"payload"`
}

func UseComtomLogger(r *gin.Engine) {

	r.Use(func(c *gin.Context) {
		startAt := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		body := c.Request.Body
		x, _ := ioutil.ReadAll(body)
		s := string(x)
		var data interface{}
		if err := json.Unmarshal([]byte(s), &data); err != nil {
			panic(err)
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(x))

		c.Next()

		endAt := time.Now()

		if raw != "" {
			path = path + "?" + raw
		}

		userId, _ := util.GetUserId(c)

		log := logBase{
			Ip:           c.ClientIP(),
			UserId:       userId,
			StartAt:      startAt,
			EndAt:        endAt,
			StatusCode:   c.Writer.Status(),
			Method:       c.Request.Method,
			Path:         path,
			Handlers:     c.HandlerNames(),
			ErrorMessage: c.Errors.Errors(),
			Payload:      data,
		}
		c.Errors.Errors()
		formated, _ := json.MarshalIndent(log, "", "  ")
		result := string(formated)
		if len(log.ErrorMessage) == 0 {
			fmt.Println(color.HiCyanString(result))
		} else {
			fmt.Println(color.HiRedString(result))
		}
	})
}

func UseLogger(r *gin.Engine) {
	r.Use(gin.Logger())
}

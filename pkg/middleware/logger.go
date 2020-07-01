package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lyouthzzz/poseidon-api-gateway/pkg/log"

	"time"
)

func Logger(conf log.Conf) gin.HandlerFunc {
	formatter := conf.Formatter
	if formatter == nil {
		formatter = log.DefaultFormatter
	}
	out := conf.OutWriter
	if out == nil {
		panic("no log outer")
	}
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		if raw != "" {
			path = path + "?" + raw
		}

		c.Next()

		latency := time.Now().Sub(start)
		param := log.FormatterParams{
			RequestID:      c.GetString(HeaderRequestId),
			Request:        c.Request,
			TimeStamp:      time.Now(),
			ClientIP:       c.ClientIP(),
			Method:         c.Request.Method,
			Path:           path,
			Latency:        latency,
			StatusCode:     c.Writer.Status(),
			InputBodySize:  int(c.Request.ContentLength),
			OutputBodySize: c.Writer.Size(),
			ErrorMessage:   c.Errors.ByType(gin.ErrorTypePrivate).String(),
		}
		_, _ = fmt.Fprintln(out, formatter(param))
	}
}

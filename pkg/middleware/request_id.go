package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strings"
)

const HeaderRequestId = "P-Request-Id"

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestId string
		if requestId = c.GetHeader(HeaderRequestId); requestId == "" {
			requestId = strings.ReplaceAll(uuid.New().String(), "-", "")
		}
		c.Request.Header.Set(HeaderRequestId, requestId)
		c.Set(HeaderRequestId, requestId)
		c.Next()
	}
}

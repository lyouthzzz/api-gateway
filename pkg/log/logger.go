package log

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Conf struct {
	Formatter Formatter
	OutWriter io.Writer
}

var DefaultConf = Conf{
	Formatter: DefaultFormatter,
	OutWriter: os.Stdout,
}

type FormatterParams struct {
	RequestID string

	Request *http.Request

	TimeStamp time.Time

	StatusCode int

	Latency time.Duration

	ClientIP string

	Method string

	Path string

	ErrorMessage string

	InputBodySize int

	OutputBodySize int
}

type Formatter func(params FormatterParams) string

var DefaultFormatter = func(param FormatterParams) string {
	return fmt.Sprintf("[Poseidon] %v | %s | %3d | %13v | %15s |%s %s %s",
		param.TimeStamp.Format(time.RFC3339),
		param.RequestID,
		param.StatusCode,
		param.Latency,
		param.ClientIP,
		param.Method,
		param.Path,
		param.ErrorMessage,
	)
}

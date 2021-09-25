package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Logger struct {
	Formatter LogFormatter
}

const (
	green             = "\033[32;1m"
	white             = "\033[37;1m"
	yellow            = "\033[33;1m"
	red               = "\033[31;1m"
	blue              = "\033[34;1m"
	magenta           = "\033[35;1m"
	cyan              = "\033[36;1m"
	reset             = "\033[0m"
	backgroundGreen   = "\033[42;1m"
	backgroundWhite   = "\033[47;1m"
	backgroundYellow  = "\033[43;1m"
	backgroundRed     = "\033[41;1m"
	backgroundBlue    = "\033[44;1m"
	backgroundMagenta = "\033[45;1m"
	backgroundCyan    = "\033[46;1m"
	backgroundReset   = "\033[0m"
)

type LogFormatter func(params LogFormatterParams) string

type LogFormatterParams struct {
	Request    *http.Request
	Timestamp  time.Time
	Method     string
	Path       string
	StatusCode int
	isTerminal bool
}

type LogResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// StatusCodeColor is the ANSI color for appropriately logging http status code to a terminal.
func (p *LogFormatterParams) StatusCodeColor() string {
	code := p.StatusCode

	switch {
	case code >= http.StatusOK && code < http.StatusMultipleChoices:
		return green
	case code >= http.StatusMultipleChoices && code < http.StatusBadRequest:
		return yellow
	case code >= http.StatusBadRequest && code < http.StatusInternalServerError:
		return magenta
	default:
		return red
	}
}

func (p *LogFormatterParams) MethodColor() string {
	method := p.Method

	switch method {
	case http.MethodGet:
		return backgroundBlue
	case http.MethodPost:
		return backgroundCyan
	case http.MethodPut:
		return backgroundYellow
	case http.MethodDelete:
		return backgroundRed
	case http.MethodPatch:
		return backgroundGreen
	case http.MethodHead:
		return backgroundMagenta
	case http.MethodOptions:
		return backgroundWhite
	default:
		return backgroundReset
	}
}

func (p *LogFormatterParams) ResetColor() string {
	return reset
}

// IsOutputColor indicates whether can colors be outputted to the log.
func (p *LogFormatterParams) IsOutputColor() bool {
	return p.isTerminal
}

func formatDefault(params LogFormatterParams) string {
	var methodColor, statusColor, resetColor string
	if params.IsOutputColor() {
		methodColor = params.MethodColor()
		statusColor = params.StatusCodeColor()
		resetColor = params.ResetColor()
	}
	requestLogFmt := fmt.Sprintf("| %s %-7s %s %s",
		methodColor, params.Method, resetColor,
		params.Path,
	)
	responseLogFmt := requestLogFmt + fmt.Sprintf(
		"| %s %v %s %s",
		statusColor, params.StatusCode, http.StatusText(params.StatusCode), resetColor,
	)
	if params.StatusCode == 0 {
		return requestLogFmt
	}
	return responseLogFmt
}

func (r *LogResponseWriter) WriteHeader(code int) {
	r.statusCode = code
	r.ResponseWriter.WriteHeader(code)
}

func LogToConsole(log *log.Logger) Adapter {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var params LogFormatterParams
			params.isTerminal = true
			params.Method = r.Method
			params.Path = r.URL.Path
			log.Println(formatDefault(params))
			lw := &LogResponseWriter{ResponseWriter: w, statusCode: 200}
			handler.ServeHTTP(lw, r)
			params.StatusCode = lw.statusCode
			log.Println(formatDefault(params))
		})
	}
}

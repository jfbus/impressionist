package handler

import (
	"net/http"
	"runtime"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/jfbus/impressionist/log"
)

// Recovery is a Negroni middleware that recovers from any panics and writes a 500 if there was one.
type Logger struct {
	StackSize int
}

// NewRecovery returns a new instance of Recovery
func NewLogger() *Logger {
	return &Logger{
		StackSize: 1024 * 8,
	}
}

func (l *Logger) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer func() {
		if err := recover(); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			stack := make([]byte, l.StackSize)
			stack = stack[:runtime.Stack(stack, false)]

			log.WithRequestId(r.Header.Get("X-Request-Id")).Errorf("%s %s - fatal error %s\n%s", r.Method, r.URL.Path, err, stack)
		}
	}()
	start := time.Now()

	next(rw, r)

	res := rw.(negroni.ResponseWriter)
	log.WithRequestId(r.Header.Get("X-Request-Id")).Infof("%s %s - %v %s in %v", r.Method, r.URL.Path, res.Status(), http.StatusText(res.Status()), time.Since(start))
}

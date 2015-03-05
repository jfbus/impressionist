package context

import (
	"errors"
	"net/http"

	"golang.org/x/net/context"
)

var ErrTimeout = errors.New("Timeout")

func New(r *http.Request) context.Context {
	return context.WithValue(context.Background(), "Request-Id", r.Header.Get("X-Request-Id"))
}

func DoWithTimeOut(ctx context.Context, fn func() error) error {
	select {
	case <-ctx.Done():
		return ErrTimeout
	default:
		return fn()
	}
	return nil
}

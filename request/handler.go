package request

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/jfbus/impressionist/config"
	ctxt "github.com/jfbus/impressionist/context"
	"github.com/jfbus/impressionist/filter"
	"github.com/jfbus/impressionist/provider"
	"github.com/jfbus/impressionist/storage"
	"golang.org/x/net/context"
)

var (
	ErrBadRequest = errors.New("bad request)")
)

func Display(w http.ResponseWriter, req *http.Request) {
	cfg := config.Get()

	ctx, cancel := context.WithTimeout(ctxt.New(req), cfg.Http.TimeOut)
	defer cancel()

	r, err := ParseRequest(ctx, req.URL.Query(), req.URL.Path)
	if err != nil {
		Error(err, w)
		return
	}
	i, err := storage.Read(r.Storage, r.File)
	if err != nil {
		Error(err, w)
		return
	}
	i, err = Work(Job{Ctx: ctx, Image: i, FilterChain: r.FilterChain})
	if err == nil {
		writer := provider.GetWriter(r.Format)
		if writer == nil {
			Error(ErrBadRequest, w)
			return
		}
		err = writer.Write(i, r.Quality, w)
	}
	if err != nil {
		Error(err, w)
	}
}

func Error(err error, w http.ResponseWriter) {
	switch err {
	case storage.ErrStorageNotFound:
		w.WriteHeader(http.StatusNotFound)
	case storage.ErrFileNotFound:
		w.WriteHeader(http.StatusNotFound)
	case storage.ErrAccessDenied:
		w.WriteHeader(http.StatusForbidden)
	case filter.ErrFilterNotFound:
		w.WriteHeader(http.StatusBadRequest)
	case filter.ErrBadFilterParameter:
		w.WriteHeader(http.StatusBadRequest)
	case filter.ErrMissingFilterParameter:
		w.WriteHeader(http.StatusBadRequest)
	case ErrBadRequest:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusServiceUnavailable)
	}
	w.Write([]byte(fmt.Sprintf("%s", err)))
}

package handler

import (
	"fmt"
	"image"
	"net/http"

	"github.com/jfbus/impressionist/action"
	"github.com/jfbus/impressionist/config"
	ctxt "github.com/jfbus/impressionist/context"
	"github.com/jfbus/impressionist/storage"
	"golang.org/x/net/context"
)

func Display(w http.ResponseWriter, req *http.Request) {
	cfg := config.Get()

	ctx, cancel := context.WithTimeout(ctxt.New(req), cfg.Http.TimeOut)
	defer cancel()

	c, writer, err := action.Build(ctx, req.URL.Query(), req.URL.Path)
	if err == nil {
		var i image.Image
		i, err = Work(Job{Ctx: ctx, ActionChain: c})
		if err == nil {
			writer.Write(i, w)
		}
	}
	if err != nil {
		switch err {
		case storage.ErrStorageNotFound:
			w.WriteHeader(http.StatusNotFound)
		case storage.ErrFileNotFound:
			w.WriteHeader(http.StatusNotFound)
		case storage.ErrAccessDenied:
			w.WriteHeader(http.StatusForbidden)
		case action.ErrBadAction:
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusServiceUnavailable)
		}
		w.Write([]byte(fmt.Sprintf("%s", err)))
	}
}

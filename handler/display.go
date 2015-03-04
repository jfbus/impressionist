package handler

import (
	"fmt"
	"net/http"

	"github.com/jfbus/impressionist/action"
	"github.com/jfbus/impressionist/storage"
)

func Display(w http.ResponseWriter, req *http.Request) {
	c, err := action.Build(req.URL.Query(), req.URL.Path)
	if err == nil {
		err = c.Apply(w)
	}
	if err != nil {
		switch err {
		case storage.ErrStorageNotFound:
			w.WriteHeader(http.StatusNotFound)
		case storage.ErrFileNotFound:
			w.WriteHeader(http.StatusNotFound)
		case storage.ErrAccessDenied:
			w.WriteHeader(http.StatusForbidden)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		w.Write([]byte(fmt.Sprintf("%s", err)))
	}
}

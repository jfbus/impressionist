package handler

import (
	"fmt"
	"net/http"

	"github.com/jfbus/impressionist/action"
)

func Display(w http.ResponseWriter, req *http.Request) {
	c, err := action.Build(req.URL.Query(), req.URL.Path)
	if err == nil {
		err = c.Apply(w)
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%s", err)))
	}
}

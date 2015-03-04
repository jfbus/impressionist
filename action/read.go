package action

import (
	"image"
	"net/http"

	"github.com/jfbus/impressionist/storage"
)

type Read struct {
	Storage string
	File    string
}

func (r *Read) Apply(i image.Image, w http.ResponseWriter) (image.Image, error) {
	i, err := storage.Read(r.Storage, r.File)
	return i, err
}

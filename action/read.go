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
	s, err := storage.Get(r.Storage)
	if err != nil {
		return nil, err
	}
	fd, err := s.Open(r.File)
	if err != nil {
		return nil, err
	}
	i, _, err = image.Decode(fd)
	fd.Close()
	return i, err
}

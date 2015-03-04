package action

import (
	"image"
	"net/http"

	"github.com/jfbus/impressionist/output"
)

type WritePNG struct{}

func (p *WritePNG) Apply(i image.Image, w http.ResponseWriter) (image.Image, error) {
	return nil, output.WriteHttpPNG(i, w)
}

type WriteJPEG struct {
	Quality int
}

func (j *WriteJPEG) Apply(i image.Image, w http.ResponseWriter) (image.Image, error) {
	return nil, output.WriteHttpJPEG(i, w, j.Quality)
}

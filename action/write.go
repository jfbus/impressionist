package action

import (
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
)

type WritePNG struct{}

func (p *WritePNG) Apply(i image.Image, w http.ResponseWriter) (image.Image, error) {
	w.Header().Set("Content-Type", "image/png")
	return nil, png.Encode(w, i)
}

type WriteJPEG struct {
	Quality int
}

func (j *WriteJPEG) Apply(i image.Image, w http.ResponseWriter) (image.Image, error) {
	w.Header().Set("Content-Type", "image/jpeg")
	return nil, jpeg.Encode(w, i, &jpeg.Options{j.Quality})
}

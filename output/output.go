package output

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
)

type Writer interface {
	Write(image.Image, io.Writer) error
	WriteHttp(image.Image, http.ResponseWriter) error
}

type PNGWriter struct{}

func (p *PNGWriter) Write(i image.Image, w io.Writer) error {
	return png.Encode(w, i)
}

func (p *PNGWriter) WriteHttp(i image.Image, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "image/png")
	return p.Write(i, w)
}

type JPEGWriter struct {
	Quality int
}

func (j *JPEGWriter) Write(i image.Image, w io.Writer) error {
	return jpeg.Encode(w, i, &jpeg.Options{j.Quality})
}

func (j *JPEGWriter) WriteHttp(i image.Image, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "image/jpeg")
	return j.Write(i, w)
}

package output

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"

	"github.com/jfbus/impressionist/config"
)

func Init(cfg config.JPEGConfig) {
	pngW = &PNGWriter{}
	jpegW = &JPEGWriter{cfg.Quality}
}

type PNGWriter struct{}

var pngW *PNGWriter

func (p *PNGWriter) Write(i image.Image, w io.Writer) error {
	return png.Encode(w, i)
}

func (p *PNGWriter) WriteHttp(i image.Image, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "image/png")
	return p.Write(i, w)
}

func WriteHttpPNG(i image.Image, w http.ResponseWriter) error {
	return pngW.WriteHttp(i, w)
}

type JPEGWriter struct {
	DefaultQuality int
}

var jpegW *JPEGWriter

func (j *JPEGWriter) Write(i image.Image, w io.Writer, quality int) error {
	if quality <= 0 || quality > 100 {
		quality = j.DefaultQuality
	}
	return jpeg.Encode(w, i, &jpeg.Options{quality})
}

func (j *JPEGWriter) WriteHttp(i image.Image, w http.ResponseWriter, quality int) error {
	w.Header().Set("Content-Type", "image/jpeg")
	return j.Write(i, w, quality)
}

func WriteHttpJPEG(i image.Image, w http.ResponseWriter, quality int) error {
	return jpegW.WriteHttp(i, w, quality)
}

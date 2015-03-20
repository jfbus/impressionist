package pure

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/jfbus/impressionist/img"
)

type PNGWriter struct{}

func (p *PNGWriter) Write(i img.Img, quality int, w io.Writer) error {
	return png.Encode(w, i.(image.Image))
}

type JPEGWriter struct{}

func (j *JPEGWriter) Write(i img.Img, quality int, w io.Writer) error {
	return jpeg.Encode(w, i.(image.Image), &jpeg.Options{quality})
}

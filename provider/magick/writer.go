package magick

import (
	"io"

	"github.com/jfbus/impressionist/img"
	"gopkgs.com/magick.v1"
)

type PNGWriter struct{}

func (p *PNGWriter) Write(i img.Img, quality int, w io.Writer) error {
	info := magick.NewInfo()
	info.SetFormat("PNG")
	info.SetQuality(quality)
	return i.(*magick.Image).Encode(w, info)
}

type JPEGWriter struct{}

func (j *JPEGWriter) Write(i img.Img, quality int, w io.Writer) error {
	info := magick.NewInfo()
	info.SetFormat("JPEG")
	info.SetQuality(quality)
	return i.(*magick.Image).Encode(w, info)
}

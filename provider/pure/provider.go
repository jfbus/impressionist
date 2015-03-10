package pure

import (
	"image"
	"io"

	"github.com/jfbus/impressionist/filter"
	"github.com/jfbus/impressionist/img"
	"github.com/jfbus/impressionist/provider"
)

type Provider struct{}

func (p *Provider) Decode(r io.Reader) (img.Img, error) {
	i, _, err := image.Decode(r)
	return i, err
}

func (p *Provider) WriterMap() map[string]provider.Writer {
	return map[string]provider.Writer{
		"png":  &PNGWriter{},
		"jpeg": &JPEGWriter{},
	}
}

func (p *Provider) FilterMap() map[string]filter.FilterBuilder {
	return map[string]filter.FilterBuilder{
		"c":  CropBuilder,
		"s":  ResizeBuilder,
		"gs": GrayscaleBuilder,
		"f":  FlipBuilder,
		"r":  RotateBuilder,
		"b":  BlurBuilder,
	}
}

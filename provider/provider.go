package provider

import (
	"io"

	"github.com/jfbus/impressionist/filter"
	"github.com/jfbus/impressionist/img"
)

type Writer interface {
	Write(i img.Img, quality int, w io.Writer) error
}

type Provider interface {
	Image() img.Img
	Decode(r io.Reader) (img.Img, error)
	FilterMap() map[string]filter.FilterBuilder
	WriterMap() map[string]Writer
}

var provider Provider

func Set(p Provider) {
	provider = p
}

func Decode(r io.Reader) (img.Img, error) {
	return provider.Decode(r)
}
func GetFilterMap() map[string]filter.FilterBuilder {
	return provider.FilterMap()
}
func GetWriter(format string) Writer {
	return provider.WriterMap()[format]
}

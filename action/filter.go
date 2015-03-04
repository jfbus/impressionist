package action

import (
	"image"
	"net/http"

	"github.com/disintegration/gift"
)

type Filter struct {
	Filter *gift.GIFT
}

func (f *Filter) Apply(i image.Image, w http.ResponseWriter) (image.Image, error) {
	dst := image.NewRGBA(f.Filter.Bounds(i.Bounds()))
	f.Filter.Draw(dst, i)
	return dst, nil
}

package action

import (
	"image"

	"github.com/disintegration/gift"
	"golang.org/x/net/context"
)

type Filter struct {
	Filter *gift.GIFT
}

func (f *Filter) Apply(ctx context.Context, i image.Image) (image.Image, error) {
	dst := image.NewRGBA(f.Filter.Bounds(i.Bounds()))
	f.Filter.Draw(dst, i)
	return dst, nil
}

package filter

import (
	cntxt "github.com/jfbus/impressionist/context"
	"github.com/jfbus/impressionist/img"
	"golang.org/x/net/context"
)

type Image interface{}

type Filter interface {
	Apply(img.Img) (img.Img, error)
}

type Chain []Filter

func (c Chain) Apply(ctx context.Context, i img.Img) (img.Img, error) {
	var err error
	for _, f := range c {
		err = cntxt.DoWithTimeOut(ctx, func() error {
			i, err = f.Apply(i)
			return err
		})
		if err != nil {
			return nil, err
		}
	}
	return i, err
}

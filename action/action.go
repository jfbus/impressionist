package action

import (
	"image"

	cntxt "github.com/jfbus/impressionist/context"
	"golang.org/x/net/context"
)

type Action interface {
	Apply(context.Context, image.Image) (image.Image, error)
}

type ActionChain []Action

func (c ActionChain) Apply(ctx context.Context) (image.Image, error) {
	i := image.Image(nil)
	err := cntxt.DoWithTimeOut(ctx, func() error {
		var err error
		for _, a := range c {
			i, err = a.Apply(ctx, i)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return i, err
}

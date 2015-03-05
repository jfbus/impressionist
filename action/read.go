package action

import (
	"image"

	cntxt "github.com/jfbus/impressionist/context"
	"github.com/jfbus/impressionist/storage"
	"golang.org/x/net/context"
)

type Read struct {
	Storage string
	File    string
}

func (r *Read) Apply(ctx context.Context, _ image.Image) (image.Image, error) {
	var i image.Image
	err := cntxt.DoWithTimeOut(ctx, func() error {
		var err error
		i, err = storage.Read(r.Storage, r.File)
		return err
	})
	return i, err
}

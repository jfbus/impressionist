package action

import (
	"image"
	"net/http"
)

type Action interface {
	Apply(image.Image, http.ResponseWriter) (image.Image, error)
}

type ActionChain []Action

func (c ActionChain) Apply(w http.ResponseWriter) error {
	i := image.Image(nil)
	var err error
	for _, a := range c {
		i, err = a.Apply(i, w)
		if err != nil {
			return err
		}
	}
	return nil
}

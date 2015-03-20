package magick

import (
	"math"

	"github.com/jfbus/impressionist/filter"
	"github.com/jfbus/impressionist/img"

	"gopkgs.com/magick.v1"
)

type Crop struct {
	rect magick.Rect
}

func (c *Crop) Apply(i img.Img) (img.Img, error) {
	return i.(*magick.Image).Crop(c.rect)
}

func CropBuilder(parts []string) (filter.Filter, int, error) {
	if len(parts) < 2 {
		return nil, 1, filter.ErrMissingFilterParameter
	}
	x, y, w, h, err := filter.ParseRect(parts[1])
	if err != nil {
		return nil, 2, err
	}
	return &Crop{magick.Rect{x, y, uint(w), uint(h)}}, 2, nil
}

type Resize struct {
	w, h int
	mod  byte
}

func (r *Resize) Apply(i img.Img) (img.Img, error) {
	w, h := r.w, r.h
	if w != 0 && h != 0 && r.mod == '+' {
		srcW := i.(*magick.Image).Width()
		srcH := i.(*magick.Image).Height()
		ratio := math.Min(float64(h)/float64(srcH), float64(w)/float64(srcW))
		w = int(math.Max(1.0, math.Floor(float64(srcW)*ratio+0.5)))
		h = int(math.Max(1.0, math.Floor(float64(srcH)*ratio+0.5)))
	}
	if r.mod == '-' {
		return i.(*magick.Image).CropResize(r.w, r.h, magick.FLanczos, magick.CSCenter)
	}
	return i.(*magick.Image).Resize(w, h, magick.FLanczos)
}

func ResizeBuilder(parts []string) (filter.Filter, int, error) {
	if len(parts) < 2 {
		return nil, 1, filter.ErrMissingFilterParameter
	}
	dim := parts[1]
	mod := dim[len(dim)-1]
	switch mod {
	case ' ':
		mod = '+'
		fallthrough
	case '+':
		fallthrough
	case '-':
		dim = dim[:len(dim)-1]
	default:
		mod = 0
	}
	w, h, err := filter.ParseDimensions(dim)
	if err != nil {
		return nil, 2, err
	}
	return &Resize{w, h, mod}, 2, nil
}

type Grayscale struct{}

func (g *Grayscale) Apply(i img.Img) (img.Img, error) {
	return i.(*magick.Image).TransformColorspace(magick.GRAY)
}

func GrayscaleBuilder(parts []string) (filter.Filter, int, error) {
	return &Grayscale{}, 1, nil
}

type Flip struct{}

func (f *Flip) Apply(i img.Img) (img.Img, error) {
	return i.(*magick.Image).Flip()
}

type Flop struct{}

func (f *Flop) Apply(i img.Img) (img.Img, error) {
	return i.(*magick.Image).Flop()
}

func FlipBuilder(parts []string) (filter.Filter, int, error) {
	if len(parts) < 2 {
		return nil, 1, filter.ErrMissingFilterParameter
	}
	switch parts[1] {
	case "h":
		return &Flop{}, 2, nil
	case "v":
		return &Flip{}, 2, nil
	}
	return nil, 2, filter.ErrBadFilterParameter
}

package pure

import (
	"errors"
	"image"
	"strconv"

	"github.com/jfbus/impressionist/filter"
	"github.com/jfbus/impressionist/img"

	"github.com/disintegration/imaging"
)

var (
	ErrFilterNotFound         = errors.New("filter not found")
	ErrBadFilterParameter     = errors.New("bad filter parameter")
	ErrMissingFilterParameter = errors.New("missing filter parameter")
)

type Crop struct {
	rect image.Rectangle
}

func (c *Crop) Apply(i img.Img) (img.Img, error) {
	return imaging.Crop(i.(image.Image), c.rect), nil
}

func CropBuilder(parts []string) (filter.Filter, int, error) {
	if len(parts) < 2 {
		return nil, 1, ErrMissingFilterParameter
	}
	x, y, w, h, err := filter.ParseRect(parts[1])
	if err != nil {
		return nil, 2, err
	}
	return &Crop{image.Rect(x, y, x+w, y+h)}, 2, nil
}

type Resize struct {
	w, h int
}

func (r *Resize) Apply(i img.Img) (img.Img, error) {
	return imaging.Resize(i.(image.Image), r.w, r.h, imaging.Lanczos), nil
}

func ResizeBuilder(parts []string) (filter.Filter, int, error) {
	if len(parts) < 2 {
		return nil, 1, ErrMissingFilterParameter
	}
	w, h, err := filter.ParseDimensions(parts[1])
	if err != nil {
		return nil, 2, err
	}
	return &Resize{w, h}, 2, nil
}

type Apply struct {
	fn func(image.Image) *image.NRGBA
}

func (a *Apply) Apply(i img.Img) (img.Img, error) {
	return a.fn(i.(image.Image)), nil
}

func GrayscaleBuilder(parts []string) (filter.Filter, int, error) {
	return &Apply{imaging.Grayscale}, 1, nil
}

func FlipBuilder(parts []string) (filter.Filter, int, error) {
	if len(parts) < 2 {
		return nil, 1, ErrMissingFilterParameter
	}
	switch parts[1] {
	case "h":
		return &Apply{imaging.FlipH}, 2, nil
	case "v":
		return &Apply{imaging.FlipV}, 2, nil
	}
	return nil, 2, ErrBadFilterParameter
}

func RotateBuilder(parts []string) (filter.Filter, int, error) {
	if len(parts) < 2 {
		return nil, 1, ErrMissingFilterParameter
	}
	switch parts[1] {
	case "90":
		return &Apply{imaging.Rotate90}, 2, nil
	case "180":
		return &Apply{imaging.Rotate180}, 2, nil
	case "270":
		return &Apply{imaging.Rotate270}, 2, nil
	}
	return nil, 2, ErrBadFilterParameter
}

type Blur struct {
	s float64
}

func (b *Blur) Apply(i img.Img) (img.Img, error) {
	return imaging.Blur(i.(image.Image), b.s), nil
}

func BlurBuilder(parts []string) (filter.Filter, int, error) {
	if len(parts) < 2 {
		return nil, 1, ErrMissingFilterParameter
	}
	s, err := strconv.ParseFloat(parts[1], 32)
	if err != nil {
		return nil, 2, ErrBadFilterParameter
	}
	return &Blur{s}, 2, nil
}

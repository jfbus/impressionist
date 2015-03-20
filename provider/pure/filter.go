package pure

import (
	"image"
	"math"
	"strconv"

	"github.com/jfbus/impressionist/filter"
	"github.com/jfbus/impressionist/img"

	"github.com/disintegration/imaging"
)

type Crop struct {
	rect image.Rectangle
}

func (c *Crop) Apply(i img.Img) (img.Img, error) {
	return imaging.Crop(i.(image.Image), c.rect), nil
}

func CropBuilder(parts []string) (filter.Filter, int, error) {
	if len(parts) < 2 {
		return nil, 1, filter.ErrMissingFilterParameter
	}
	x, y, w, h, err := filter.ParseRect(parts[1])
	if err != nil {
		return nil, 2, err
	}
	return &Crop{image.Rect(x, y, x+w, y+h)}, 2, nil
}

type Resize struct {
	w, h int
	mod  byte
}

func (r *Resize) Apply(i img.Img) (img.Img, error) {
	w, h, crop := r.w, r.h, false
	if w != 0 && h != 0 && r.mod != 0 {
		srcW := i.(image.Image).Bounds().Max.X
		srcH := i.(image.Image).Bounds().Max.Y
		var ratio float64
		switch r.mod {
		case '+':
			ratio = math.Min(float64(h)/float64(srcH), float64(w)/float64(srcW))
		case '-':
			ratio = math.Max(float64(h)/float64(srcH), float64(w)/float64(srcW))
			crop = true
		}
		w = int(math.Max(1.0, math.Floor(float64(srcW)*ratio+0.5)))
		h = int(math.Max(1.0, math.Floor(float64(srcH)*ratio+0.5)))
	}
	n := imaging.Resize(i.(image.Image), w, h, imaging.Lanczos)
	if crop {
		tmpW := n.Bounds().Max.X
		tmpH := n.Bounds().Max.Y
		rect := image.Rectangle{image.Point{(tmpW - r.w) / 2, (tmpH - r.h) / 2}, image.Point{(tmpW + r.w) / 2, (tmpH + r.h) / 2}}
		return imaging.Crop(n, rect), nil
	}
	return n, nil
}

func ResizeBuilder(parts []string) (filter.Filter, int, error) {
	if len(parts) < 2 {
		return nil, 1, filter.ErrMissingFilterParameter
	}
	mod, dim := filter.ParseModifier(parts[1], " +-")
	// just in case + is interpreted as a space...
	if mod == ' ' {
		mod = '+'
	}
	w, h, err := filter.ParseDimensions(dim)
	if err != nil {
		return nil, 2, err
	}
	return &Resize{w, h, mod}, 2, nil
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
		return nil, 1, filter.ErrMissingFilterParameter
	}
	switch parts[1] {
	case "h":
		return &Apply{imaging.FlipH}, 2, nil
	case "v":
		return &Apply{imaging.FlipV}, 2, nil
	}
	return nil, 2, filter.ErrBadFilterParameter
}

func RotateBuilder(parts []string) (filter.Filter, int, error) {
	if len(parts) < 2 {
		return nil, 1, filter.ErrMissingFilterParameter
	}
	switch parts[1] {
	case "90":
		return &Apply{imaging.Rotate90}, 2, nil
	case "180":
		return &Apply{imaging.Rotate180}, 2, nil
	case "270":
		return &Apply{imaging.Rotate270}, 2, nil
	}
	return nil, 2, filter.ErrBadFilterParameter
}

type Blur struct {
	s float64
}

func (b *Blur) Apply(i img.Img) (img.Img, error) {
	return imaging.Blur(i.(image.Image), b.s), nil
}

func BlurBuilder(parts []string) (filter.Filter, int, error) {
	if len(parts) < 2 {
		return nil, 1, filter.ErrMissingFilterParameter
	}
	s, err := strconv.ParseFloat(parts[1], 32)
	if err != nil {
		return nil, 2, filter.ErrBadFilterParameter
	}
	return &Blur{s}, 2, nil
}

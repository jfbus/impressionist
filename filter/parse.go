package filter

import (
	"errors"
	"image"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"

	"github.com/disintegration/gift"
)

var (
	ErrFilterNotFound         = errors.New("filter not found")
	ErrBadFilterParameter     = errors.New("bad filter parameter")
	ErrMissingFilterParameter = errors.New("missing filter parameter")
)

type parseFn func([]string) (gift.Filter, int, error)

func Parse(str string) (*gift.GIFT, error) {
	// predefined filter
	g := predefined(str)
	if g != nil {
		return g, nil
	}
	// parse filter
	g = gift.New()
	parts := strings.Split(str, ",")
	for {
		fn, err := parseFilter(parts[0])
		if err != nil {
			return nil, err
		}
		f, adv, err := fn(parts)
		if err != nil {
			return nil, err
		}
		g.Add(f)
		if len(parts) <= adv {
			break
		}
		parts = parts[adv:]
	}
	return g, nil
}

func parseFilter(code string) (parseFn, error) {
	switch code {
	case "c":
		return parseCrop, nil
	case "s":
		return parseResize, nil
	case "gs":
		return parseGrayscale, nil
	case "f":
		return parseFlip, nil
	case "r":
		return parseRotate, nil
	case "b":
		return parseBlur, nil
	}
	return nil, ErrFilterNotFound
}

func parseCrop(parts []string) (gift.Filter, int, error) {
	if len(parts) < 2 {
		return nil, 1, ErrMissingFilterParameter
	}
	r, err := parseRect(parts[1])
	if err != nil {
		return nil, 2, err
	}
	return gift.Crop(r), 2, nil
}

func parseResize(parts []string) (gift.Filter, int, error) {
	if len(parts) < 2 {
		return nil, 1, ErrMissingFilterParameter
	}
	w, h, err := parseDimensions(parts[1])
	if err != nil {
		return nil, 2, err
	}
	return gift.Resize(w, h, gift.LanczosResampling), 2, nil
}

func parseGrayscale(parts []string) (gift.Filter, int, error) {
	return gift.Grayscale(), 1, nil
}

func parseFlip(parts []string) (gift.Filter, int, error) {
	if len(parts) < 2 {
		return nil, 1, ErrMissingFilterParameter
	}
	switch parts[1] {
	case "h":
		return gift.FlipHorizontal(), 2, nil
	case "v":
		return gift.FlipVertical(), 2, nil
	}
	return nil, 2, ErrBadFilterParameter
}

func parseRotate(parts []string) (gift.Filter, int, error) {
	if len(parts) < 2 {
		return nil, 1, ErrMissingFilterParameter
	}
	switch parts[1] {
	case "90":
		return gift.Rotate90(), 2, nil
	case "180":
		return gift.Rotate180(), 2, nil
	case "270":
		return gift.Rotate270(), 2, nil
	}
	return nil, 2, ErrBadFilterParameter
}

func parseBlur(parts []string) (gift.Filter, int, error) {
	if len(parts) < 2 {
		return nil, 1, ErrMissingFilterParameter
	}
	s, err := strconv.ParseFloat(parts[1], 32)
	if err != nil {
		return nil, 2, ErrBadFilterParameter
	}
	return gift.GaussianBlur(float32(s)), 2, nil
}

func parseRect(str string) (image.Rectangle, error) {
	parts := strings.Split(str, "-")
	if len(parts) != 2 {
		log.Warnf("Unable to parse rectangle %s", str)
		return image.Rectangle{}, ErrBadFilterParameter
	}
	x, y, err := parseDimensions(parts[0])
	if err != nil {
		return image.Rectangle{}, ErrBadFilterParameter
	}
	w, h, err := parseDimensions(parts[1])
	if err != nil {
		return image.Rectangle{}, ErrBadFilterParameter
	}
	return image.Rect(x, y, x+w, y+h), nil
}

func parseDimensions(str string) (int, int, error) {
	parts := strings.Split(str, "x")
	if len(parts) != 2 {
		log.Warnf("Unable to parse dimensions %s", str)
		return 0, 0, ErrBadFilterParameter
	}
	x, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		log.Warnf("Unable to parse dimensions %s (not an integer)", str)
		return 0, 0, ErrBadFilterParameter
	}
	y, err := strconv.ParseInt(parts[1], 10, 32)
	if err != nil {
		log.Warnf("Unable to parse dimensions %s (not an integer)", str)
		return 0, 0, ErrBadFilterParameter
	}
	return int(x), int(y), nil
}

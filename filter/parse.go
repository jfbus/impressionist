package filter

import (
	"errors"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
)

var (
	ErrFilterNotFound         = errors.New("filter not found")
	ErrBadFilterParameter     = errors.New("bad filter parameter")
	ErrMissingFilterParameter = errors.New("missing filter parameter")
)

type FilterBuilder func(parts []string) (Filter, int, error)

func Parse(ctx context.Context, m map[string]FilterBuilder, str string) (Chain, error) {
	// predefined filter
	c := predefined(str)
	if c != nil {
		return c, nil
	}
	// parse filter
	c = Chain{}
	parts := strings.Split(str, ",")
	for {
		fn, found := m[parts[0]]
		if !found {
			return nil, ErrFilterNotFound
		}
		f, adv, err := fn(parts)
		if err != nil {
			return nil, err
		}
		c = append(c, f)
		if len(parts) <= adv {
			break
		}
		parts = parts[adv:]
	}
	return c, nil
}

func ParseRect(str string) (int, int, int, int, error) {
	parts := strings.Split(str, "-")
	if len(parts) != 2 {
		log.Warnf("Unable to parse rectangle %s", str)
		return 0, 0, 0, 0, ErrBadFilterParameter
	}
	x, y, err := ParseDimensions(parts[0])
	if err != nil {
		return 0, 0, 0, 0, ErrBadFilterParameter
	}
	w, h, err := ParseDimensions(parts[1])
	if err != nil {
		return 0, 0, 0, 0, ErrBadFilterParameter
	}
	return x, y, w, h, nil
}

func ParseDimensions(str string) (int, int, error) {
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

package action

import (
	"errors"
	"net/url"
	"path"
	"strconv"
	"strings"

	"golang.org/x/net/context"

	"github.com/jfbus/impressionist/config"
	"github.com/jfbus/impressionist/filter"
	"github.com/jfbus/impressionist/log"
	"github.com/jfbus/impressionist/output"
)

var (
	ErrBadAction = errors.New("bad action url")
)

func Build(ctx context.Context, q url.Values, url string) (ActionChain, output.Writer, error) {
	var w output.Writer
	fmt := strings.Split(q.Get(":format"), ",")
	switch fmt[0] {
	case "png":
		w = &output.PNGWriter{}
	case "jpeg":
		q := 0
		if len(fmt) > 1 {
			qq, err := strconv.ParseInt(fmt[1], 10, 32)
			if err == nil {
				q = int(qq)
			} else {
				log.WithContext(ctx).Infof("JPEG quality %s is not a number, using default", fmt[1])
			}
			if q <= 0 || q > 100 {
				cfg := config.Get()
				q = cfg.JPEG.Quality
			}
		}
		w = &output.JPEGWriter{q}
	default:
		log.WithContext(ctx).Warnf("unknown format %s", q.Get(":format"))
		return nil, nil, ErrBadAction
	}
	f, err := filter.Parse(q.Get(":filter"))
	if err != nil {
		return nil, nil, ErrBadAction
	}
	parts := strings.Split(url, "/")
	p := ""
	for i := 3; i < len(parts)-1; i++ {
		if parts[i] == q.Get(":storage") {
			p = path.Join(parts[i+1:]...)
		}
	}
	if p == "" {
		log.WithContext(ctx).Warn("missing path")
		return nil, nil, ErrBadAction
	}
	return ActionChain{
		&Read{
			Storage: q.Get(":storage"),
			File:    p,
		},
		&Filter{f},
	}, w, nil
}

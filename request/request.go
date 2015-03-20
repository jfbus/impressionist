package request

import (
	"net/url"
	"path"
	"strconv"
	"strings"

	"github.com/jfbus/impressionist/config"
	"github.com/jfbus/impressionist/filter"
	"github.com/jfbus/impressionist/log"
	"github.com/jfbus/impressionist/provider"
	"golang.org/x/net/context"
)

type Request struct {
	FilterChain filter.Chain
	Storage     string
	File        string
	Format      string
	Quality     int
}

func ParseRequest(ctx context.Context, q url.Values, url string) (*Request, error) {
	r := Request{}

	// format
	fmt := strings.Split(q.Get(":format"), ",")
	r.Format = fmt[0]
	if len(fmt) > 1 {
		qq, err := strconv.ParseInt(fmt[1], 10, 32)
		if err == nil {
			r.Quality = int(qq)
		} else {
			log.WithContext(ctx).Infof("Image quality %s is not a number, using default", fmt[1])
		}
	}
	if r.Quality <= 0 || r.Quality > 100 {
		cfg := config.Get()
		r.Quality = cfg.Image.Quality
	}

	// filter
	var err error
	r.FilterChain, err = filter.Parse(ctx, provider.GetFilterMap(), q.Get(":filter"))
	if err != nil {
		return nil, ErrBadRequest
	}

	// file
	parts := strings.Split(url, "/")
	p := ""
	for i := 3; i < len(parts)-1; i++ {
		if parts[i] == q.Get(":storage") {
			p = path.Join(parts[i+1:]...)
		}
	}
	if p == "" {
		log.WithContext(ctx).Warn("missing path")
		return nil, ErrBadRequest
	}
	r.Storage = q.Get(":storage")
	r.File = p

	return &r, nil
}

package action

import (
	"errors"
	"net/url"
	"path"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"

	"github.com/jfbus/impressionist/filter"
)

var (
	ErrBadAction = errors.New("bad action url")
)

func Build(q url.Values, url string) (ActionChain, error) {
	var w Action
	fmt := strings.Split(q.Get(":format"), ",")
	switch fmt[0] {
	case "png":
		w = &WritePNG{}
	case "jpeg":
		q := 80
		if len(fmt) > 1 {
			qq, err := strconv.ParseInt(fmt[1], 10, 32)
			if err == nil {
				q = int(qq)
			} else {
				log.Infof("JPEG quality %s is not number, ignored", qq)
			}
		}
		w = &WriteJPEG{q}
	default:
		log.Warnf("unknown format %s", q.Get(":format"))
		return nil, ErrBadAction
	}
	f, err := filter.Parse(q.Get(":filter"))
	if err != nil {
		return nil, ErrBadAction
	}
	parts := strings.Split(url, "/")
	p := ""
	for i := 3; i < len(parts)-1; i++ {
		if parts[i] == q.Get(":storage") {
			p = path.Join(parts[i+1:]...)
		}
	}
	if p == "" {
		log.Warn("missing path")
		return nil, ErrBadAction
	}
	return ActionChain{
		&Read{
			Storage: q.Get(":storage"),
			File:    p,
		},
		&Filter{f},
		w,
	}, nil
}

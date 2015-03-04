package filter

import (
	"github.com/disintegration/gift"
	"github.com/jfbus/impressionist/config"
)

var predef = map[string]*gift.GIFT{}

func Init(cfg []config.FilterConfig) {
	var err error
	for _, f := range cfg {
		predef[f.Name], err = Parse(f.Definition)
		if err != nil {
			panic("bad filter definition for " + f.Name)
		}
	}
}

func predefined(str string) *gift.GIFT {
	return predef[str]
}

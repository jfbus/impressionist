package filter

import (
	"github.com/jfbus/impressionist/config"
	"golang.org/x/net/context"
)

var predef = map[string]Chain{}

func Init(ctx context.Context, cfg []config.FilterConfig, m map[string]FilterBuilder) {
	var err error
	for _, f := range cfg {
		predef[f.Name], err = Parse(ctx, m, f.Definition)
		if err != nil {
			panic("bad filter definition for " + f.Name)
		}
	}
}

func predefined(str string) Chain {
	return predef[str]
}

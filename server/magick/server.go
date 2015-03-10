package main

import (
	"flag"
	"fmt"
	"net/http"
	"runtime"

	"github.com/bmizerany/pat"
	"github.com/codegangsta/negroni"
	"github.com/jfbus/impressionist/config"
	"github.com/jfbus/impressionist/filter"
	"github.com/jfbus/impressionist/log"
	"github.com/jfbus/impressionist/provider"
	"github.com/jfbus/impressionist/provider/magick"
	"github.com/jfbus/impressionist/request"
	"github.com/jfbus/impressionist/storage"
	"github.com/pilu/xrequestid"
	"golang.org/x/net/context"
)

func main() {

	file := flag.String("cfg", "./impressionist.json", "config file")
	debug := flag.Bool("debug", false, "debug mode")
	flag.Parse()

	cfg := config.Load(*file)
	ctx := context.Background()
	provider.Set(&magick.Provider{})
	storage.Init(cfg.Storages, cfg.Cache.Source)
	filter.Init(ctx, cfg.Filters, provider.GetFilterMap())
	request.InitWorkers(cfg.Http.Workers)
	log.Init(*debug)

	runtime.GOMAXPROCS(runtime.NumCPU())

	m := pat.New()
	m.Get(cfg.Http.Root+"/:filter/:format/:storage/", http.HandlerFunc(request.Display))

	n := negroni.New(xrequestid.New(16), request.NewLogger())
	n.UseHandler(m)
	n.Run(fmt.Sprintf(":%d", cfg.Http.Port))
}

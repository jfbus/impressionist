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
	"github.com/jfbus/impressionist/handler"
	"github.com/jfbus/impressionist/log"
	"github.com/jfbus/impressionist/storage"
	"github.com/pilu/xrequestid"
)

func main() {

	file := flag.String("cfg", "./impressionist.json", "config file")
	debug := flag.Bool("debug", false, "debug mode")
	flag.Parse()

	cfg := config.Load(*file)
	storage.Init(cfg.Storages, cfg.Cache.Source)
	filter.Init(cfg.Filters)
	handler.InitWorkers(cfg.Http.Workers)
	log.Init(*debug)

	runtime.GOMAXPROCS(runtime.NumCPU())

	m := pat.New()
	m.Get(cfg.Http.Root+"/:filter/:format/:storage/", http.HandlerFunc(handler.Display))

	n := negroni.New(xrequestid.New(16), handler.NewLogger())
	n.UseHandler(m)
	n.Run(fmt.Sprintf(":%d", cfg.Http.Port))
}

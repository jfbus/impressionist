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
	"github.com/jfbus/impressionist/storage"
)

func main() {

	file := flag.String("cfg", "./impressionist.json", "config file")
	flag.Parse()
	cfg := config.Load(*file)
	storage.Init(cfg.Storages)
	filter.Init(cfg.Filters)
	runtime.GOMAXPROCS(runtime.NumCPU())

	m := pat.New()
	m.Get(cfg.Http.Root+"/:filter/:format/:storage/", http.HandlerFunc(handler.Display))

	n := negroni.New(negroni.NewRecovery(), negroni.NewLogger())
	n.UseHandler(m)
	n.Run(fmt.Sprintf(":%d", cfg.Http.Port))
}

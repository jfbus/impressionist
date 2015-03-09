package request

import (
	ctxt "github.com/jfbus/impressionist/context"
	"github.com/jfbus/impressionist/filter"
	"github.com/jfbus/impressionist/img"
	"github.com/jfbus/impressionist/log"
	"golang.org/x/net/context"
)

type Job struct {
	Ctx         context.Context
	Image       img.Img
	FilterChain filter.Chain
	res         chan JobResponse
}

type JobResponse struct {
	i   img.Img
	err error
}

var queue chan Job

func InitWorkers(n int) {
	log.Infof("Starting %d workers", n)
	queue = make(chan Job)
	for i := 0; i < n; i++ {
		go work(queue)
	}
}

func work(queue chan Job) {
	for j := range queue {
		i, err := j.FilterChain.Apply(j.Ctx, j.Image)
		j.res <- JobResponse{i, err}
	}
}

func Work(j Job) (img.Img, error) {
	j.res = make(chan JobResponse)
	select {
	case <-j.Ctx.Done():
		return nil, ctxt.ErrTimeout
	case queue <- j:
	}
	r := <-j.res
	return r.i, r.err
}

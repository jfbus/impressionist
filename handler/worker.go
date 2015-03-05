package handler

import (
	"image"

	"github.com/jfbus/impressionist/action"
	"github.com/jfbus/impressionist/log"
	"golang.org/x/net/context"
)

type Job struct {
	Ctx         context.Context
	ActionChain action.ActionChain
	res         chan JobResponse
}

type JobResponse struct {
	i   image.Image
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
	for {
		j := <-queue
		i, err := j.ActionChain.Apply(j.Ctx)
		j.res <- JobResponse{i, err}
	}
}

func Work(j Job) (image.Image, error) {
	j.res = make(chan JobResponse)
	queue <- j
	r := <-j.res
	return r.i, r.err
}

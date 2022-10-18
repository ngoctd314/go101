package workerpool

import (
	"fmt"
	"log"
)

// worker represents the worker that execute the Job
type worker struct {
	workerPool chan chan Job // number of worker
	jobChannel chan Job      // current job
}

func newWorker(workerPool chan chan Job) *worker {
	return &worker{
		workerPool: workerPool,
		jobChannel: make(chan Job),
	}
}

func (w *worker) exec() {
	// each worker has it own goroutine
	go func() {
		for {
			// send jobChannel to workerPool
			w.workerPool <- w.jobChannel
			select {
			// case receive job from job channel
			case job := <-w.jobChannel:
				if err := job.Do(); err != nil {
					log.Println(fmt.Errorf("exec job error: %v", err))
				}
			}
		}
	}()
}

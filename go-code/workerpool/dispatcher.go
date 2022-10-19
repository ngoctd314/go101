package workerpool

type dispatcher struct {
	// A pool of workers channels that are registered with dispatcher
	workerPool chan chan Job
}

func newDispatcher(maxWorkers int) *dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &dispatcher{
		workerPool: pool,
	}
}

func (d *dispatcher) initWorker() {
	// starting n workers
	// n workers reference the same worker pool
	for i := 0; i < cap(d.workerPool); i++ {
		worker := newWorker(d.workerPool)
		worker.exec()
	}

	go d.dispatch()

}

func (d *dispatcher) dispatch() {
	for {
		select {
		// receive job from jobQueue
		case job := <-jobQueue:
			// assign job to idle worker
			worker := <-d.workerPool
			worker <- job
		}
	}
}

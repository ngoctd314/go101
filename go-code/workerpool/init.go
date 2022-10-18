package workerpool

import (
	"os"
	"strconv"
)

func init() {
	jobQueue = make(chan Job)

	maxWorker := os.Getenv("MAX_WORKER")
	mw, err := strconv.Atoi(maxWorker)
	if err != nil {
		mw = 1200
	}

	dispatcher := newDispatcher(mw)
	dispatcher.initWorker()
	go dispatcher.dispatch()
}

package workerpool

import (
	"os"
	"strconv"
)

func init() {
	maxQueue := os.Getenv("MAX_QUEUE")
	mq, err := strconv.Atoi(maxQueue)
	if err != nil {
		mq = 3000
	}

	jobQueue = make(chan Job, mq)

	maxWorker := os.Getenv("MAX_WORKER")
	mw, err := strconv.Atoi(maxWorker)
	if err != nil {
		mw = 1000
	}

	dispatcher := newDispatcher(mw)
	dispatcher.initWorker()
}

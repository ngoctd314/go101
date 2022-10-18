package workerpool

// Job represents job will run by worker pool
type Job interface {
	Do() error
}

// jobQueue is a buffered channel for enqueue and dequeue Job
var jobQueue chan Job

// EnQueue send job to jobQueue
func EnQueue(job Job) {
	jobQueue <- job
}

// DeQueue return job from jobQueue
func DeQueue() Job {
	return <-jobQueue
}

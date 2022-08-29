package cronjob

// Job interface
type Job interface {
	Run()
}

// Cron keeps track of any number of entries
type Cron struct {
	entries []any
}

// Entry consists of a schedule and the func to execute on the schedule
type Entry struct{}

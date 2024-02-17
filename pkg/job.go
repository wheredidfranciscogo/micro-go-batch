package pkg

// Job represents an individual task
type Job struct {
	Data interface{}
}

// JobResult represents the result of a job processing
type JobResult struct {
	Data  interface{}
	Error error
}

// testing purpouse creates a new JoB
func NewJob(data interface{}) *Job {
	return &Job{Data: data}
}

// creates a new JobResult instance with the given data and error
func NewJobResult(data interface{}, err error) *JobResult {
	return &JobResult{Data: data, Error: err}
}
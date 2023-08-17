package main

type JobPool struct {
	Jobs []Job `json:"jobs"`
}

func NewJobPool() *JobPool {
	return &JobPool{Jobs: []Job{}}
}

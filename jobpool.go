package main

type JobPool struct {
	Jobs []Job
}

func NewJobPool() *JobPool {
	return &JobPool{Jobs: []Job{}}
}

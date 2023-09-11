package main

import "fmt"

type Job struct {
	Handle ExecHandle `json:"handle"`
}

func NewJob() Job {
	return Job{Handle: ExecHandle(0)}
}

func (j *Job) Run() {
	fmt.Printf("Job<%v> start\n", j.Handle)
}

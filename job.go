package main

import "fmt"

type Job struct {
	Id uint32 `json:"id"`
}

func NewJob() *Job {
	return &Job{Id: 0}
}

func (j *Job) Run() {
	fmt.Printf("Job<%d> start\n", j.Id)
}

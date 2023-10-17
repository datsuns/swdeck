package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
)

type Job struct {
	Name    string `json:"name"`
	Command string `json:"command"`
	Type    string `json:"type"` // "app" or "mp3"
}

type JobPool struct {
	Jobs []Job `json:"jobs"`
}

func (jp *JobPool) Load() error {
	data, err := ioutil.ReadFile("jobs.json")
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &jp)
}

func (jp *JobPool) Save() error {
	data, err := json.Marshal(jp)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("jobs.json", data, 0644)
}

func (jp *JobPool) Add(job Job) {
	jp.Jobs = append(jp.Jobs, job)
	jp.Save()
}

func (jp *JobPool) Delete(index int) {
	if index >= 0 && index < len(jp.Jobs) {
		jp.Jobs = append(jp.Jobs[:index], jp.Jobs[index+1:]...)
		jp.Save()
	}
}

func (jp *JobPool) Get() []Job {
	return jp.Jobs
}

func (jp *JobPool) Launch(job Job) error {
	switch strings.ToLower(job.Type) {
	case "app":
		return exec.Command(job.Command).Start()
	case "mp3":
		return playMP3(job.Command)
	default:
		return errors.New("unknown job type")
	}
}

func playMP3(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}

	c, ready, err := oto.NewContext(d.SampleRate(), 2, 2)
	if err != nil {
		return err
	}
	<-ready

	p := c.NewPlayer(d)
	defer p.Close()
	p.Play()

	fmt.Printf("Length: %d[bytes]\n", d.Length())
	for {
		time.Sleep(time.Second)
		if !p.IsPlaying() {
			break
		}
	}

	return nil
}

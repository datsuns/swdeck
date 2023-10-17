package main

import (
	"context"
)

// App struct
type App struct {
	ctx  context.Context
	Pool *JobPool
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{Pool: &JobPool{}}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.Pool.Load()
}

func (a *App) Add(j Job) {
	a.Pool.Add(j)
}

func (a *App) Launch(j Job) {
	a.Pool.Launch(j)
}

func (a *App) Get() []Job {
	return a.Pool.Get()
}

func (a *App) Delete(index int) {
	a.Pool.Delete(index)
}

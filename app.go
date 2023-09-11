package main

import (
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx  context.Context
	Pool *JobPool
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{Pool: NewJobPool()}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Register(e Entry) ExecHandle {
	fmt.Println("go: register ", e.Type)
	return ExecHandle(1)
}

func (a *App) Run(h ExecHandle) {
	fmt.Println("go: run ", h)
}

func (a *App) Load(h ExecHandle) Job {
	ret := NewJob()
	fmt.Println("go: Load ", h, " -> ", ret)
	return ret
}

func (a *App) LoadAll() []Job {
	ret := []Job{NewJob()}
	fmt.Println("go: LoadAll ", ret[0].Handle)
	return ret
}

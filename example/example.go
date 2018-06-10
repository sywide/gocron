package main

import (
	"fmt"

	"github.com/sywide/gocron"
)

func task() {
	fmt.Println("I am runnning task.")
}

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}

var (
	functions = map[string]interface{}{
		"task":           task,
		"taskWithParams": taskWithParams,
	}
)

func main() {

	// Bind all existing functions
	funcs := gocron.FuncBindAll(functions)

	// Do jobs with params
	gocron.Every(1).Second().Do(funcs.List["taskWithParams"], 1, "hello")

	// Do jobs without params
	gocron.Every(1).Second().Do(funcs.List["task"])
	gocron.Every(2).Seconds().Do(funcs.List["task"])
	gocron.Every(1).Minute().Do(funcs.List["task"])
	gocron.Every(2).Minutes().Do(funcs.List["task"])
	gocron.Every(1).Hour().Do(funcs.List["task"])
	gocron.Every(2).Hours().Do(funcs.List["task"])
	gocron.Every(1).Day().Do(funcs.List["task"])
	gocron.Every(2).Days().Do(funcs.List["task"])

	// Do jobs on specific weekday
	gocron.Every(1).Monday().Do(funcs.List["task"])
	gocron.Every(1).Thursday().Do(funcs.List["task"])

	// function At() take a string like 'hour:min'
	gocron.Every(1).Day().At("10:30").Do(funcs.List["task"])
	gocron.Every(1).Monday().At("18:30").Do(funcs.List["task"])

	// remove, clear and next_run
	_, time := gocron.NextRun()
	fmt.Println(time)

	// gocron.Remove(task)
	// gocron.Clear()

	// function Start start all the pending jobs
	<-gocron.Start()

	// also , you can create a your new scheduler,
	// to run two scheduler concurrently
	s := gocron.NewScheduler()
	s.Every(3).Seconds().Do(funcs.List["task"])
	<-s.Start()
}

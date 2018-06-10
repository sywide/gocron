# goCron: A Golang Job Scheduling Package

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](http://godoc.org/github.com/jasonlvhit/gocron)
[![Stories in Ready](https://badge.waffle.io/jasonlvhit/gocron.png?label=ready&title=Ready)](https://waffle.io/jasonlvhit/gocron)

goCron is a Golang job scheduling package which lets you run Go functions periodically at pre-determined interval using a simple, human-friendly syntax.

goCron is a Golang implementation of Ruby module [clockwork](<https://github.com/tomykaira/clockwork>) and Python job scheduling package [schedule](<https://github.com/dbader/schedule>), and personally, this package is my first Golang program, just for fun and practice.

See also this two great articles:

* [Rethinking Cron](http://adam.herokuapp.com/past/2010/4/13/rethinking_cron/)
* [Replace Cron with Clockwork](http://adam.herokuapp.com/past/2010/6/30/replace_cron_with_clockwork/)

**Fork by sywide**:  Add Functions bindings to use with external configuration

Back to this package, you could just use this simple API as below, to run a cron scheduler.

``` go
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
```

and full test cases and [document](http://godoc.org/github.com/jasonlvhit/gocron) will be coming soon.

Once again, thanks to the great works of Ruby clockwork and Python schedule package. BSD license is used, see the file License for detail.

Have fun!

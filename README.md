# Go Actor System

This package provides an implementation of the Actor Model in Go. It allows for building concurrent applications by using lightweight actors that communicate via message passing.

## Overview

The Actor System consists of the following components:

- **Actor System Interface**: The entry point for submitting tasks to the Actor System.
- **Task Queue**: A queue that holds the tasks submitted by the Actor System Interface.
- **Assigner Actor**: Responsible for reading tasks from the Task Queue and assigning them to available Actor instances.
- **Actors**: Lightweight entities that execute the assigned tasks concurrently. Each Actor has its own Task Queue.
- **Auto Scaler**: Dynamically increases or decreases the number of Actor instances based on the Task Queue size.

## Usage

```go
package main

import (
	"sync"
	"time"

	"github.com/forkbikash/aktor/actor"
	"github.com/forkbikash/aktor/actor_system"
	"github.com/forkbikash/aktor/entities"
)

type task struct{}

func (t *task) Execute() {}

func taskFunc() entities.Task {
	return &task{}
}

func main() {
	actorSystem := actor_system.CreateActorSystem("actor_system", &actor.Config{
		MinActor: 10,
		MaxActor: 100,
		AutoScale: actor.AutoScale{
			UpscaleQueueSize:   100,
			DownscaleQueueSize: 10,
		},
	})

	for i := 0; i < 1000; i += 1 {
		actorSystem.SubmitTask(taskFunc())
		<-time.After(2 * time.Millisecond)
	}

	systems := []*actor_system.ActorSystem{actorSystem}
	wg := &sync.WaitGroup{}
	wg.Add(len(systems))
	for _, system := range systems {
		go system.Shutdown(wg)
	}
	wg.Wait()
}
```

## What's next

- Assign task id to each task and functionality to wait for the completion of a task with specific task id
- Support for distributed environment

package actor_system

import (
	"math/rand"
	"time"

	"github.com/forkbikash/aktor/entities"
)

// SimIOTask this is task to simulate IO operation.
// It would randomly take [0-10) milliseconds to finish
type SimIOTask struct {
	num int
}

func CreateNumberPrinterTask(i int) entities.Task {
	return &SimIOTask{
		num: i,
	}
}

func (t *SimIOTask) Execute() {
	// Simulate IO operation taking [0-10) milliseconds as documented
	duration := time.Duration(rand.Intn(10)) * time.Millisecond
	<-time.After(duration)
}

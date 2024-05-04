package actor

import (
	"sync"

	"github.com/forkbikash/aktor/entities"
)

type TaskActorPool struct {
	pool     []entities.Actor
	poolLock *sync.Mutex
	wg       *sync.WaitGroup
}

func CreateTaskActorPool(wg *sync.WaitGroup) *TaskActorPool {
	return &TaskActorPool{
		pool:     []entities.Actor{},
		poolLock: &sync.Mutex{},
		wg:       wg,
	}
}

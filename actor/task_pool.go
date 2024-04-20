package actor

import (
	"aktor/entities"
	"sync"
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

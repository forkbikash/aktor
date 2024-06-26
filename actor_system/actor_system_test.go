package actor_system

import (
	"sync"
	"testing"
	"time"

	"github.com/forkbikash/aktor/actor"

	"github.com/ian-kent/go-log/log"
)

func TestIOSimulationSystem(t *testing.T) {
	ioSimSystem := CreateActorSystem("io_sim", &actor.Config{
		MinActor: 10,
		MaxActor: 100,
		AutoScale: actor.AutoScale{
			UpscaleQueueSize:   100,
			DownscaleQueueSize: 10,
		},
	})

	for i := 0; i < 1000; i += 1 {
		ioSimSystem.SubmitTask(CreateNumberPrinterTask(i))
		<-time.After(2 * time.Millisecond)
	}
	shutdown([]*ActorSystem{ioSimSystem})

}

func shutdown(systems []*ActorSystem) {

	wg := &sync.WaitGroup{}
	wg.Add(len(systems))
	for _, system := range systems {
		go system.Shutdown(wg)
	}
	wg.Wait()
	log.Debug("shutting down")

}

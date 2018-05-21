package reminder_test

import (
	"sync"
	"testing"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/persistence"
	reminder "github.com/artyomturkin/protoactor-go-reminder"
)

type testActor struct {
	reminder.Mixin

	wg *sync.WaitGroup
}

func (r *testActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *Setup:
		r.RemindMe("hello", 1*time.Millisecond, false)
	case *reminder.Remind:
		if msg.Name == "hello" {
			r.wg.Done()
		}
	}
}

func testActorProducer(wg *sync.WaitGroup) func() actor.Actor {
	return func() actor.Actor {
		return &testActor{wg: wg}
	}
}

type Setup struct{}

func TestPlugin(t *testing.T) {
	remProps := actor.FromProducer(reminder.Producer(10 * time.Millisecond)).
		WithMiddleware(persistence.Using(newProvider(5)))
	rem, err := actor.SpawnNamed(remProps, "reminder")
	if err != nil {
		t.Fatalf("failed to spawn reminder: %v", err)
	}

	wg := &sync.WaitGroup{}
	recProps := actor.FromProducer(testActorProducer(wg)).
		WithMiddleware(reminder.Middleware(rem))
	rec, err := actor.SpawnNamed(recProps, "receiver")
	if err != nil {
		t.Fatalf("failed to spawn receiver: %v", err)
	}

	rems := 10
	wg.Add(rems)
	for i := 0; i < rems; i++ {
		rec.Tell(&Setup{})
	}
	wg.Wait()

	rem.GracefulPoison()
	rec.GracefulPoison()
}

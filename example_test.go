package reminder_test

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/persistence"
	reminder "github.com/artyomturkin/protoactor-go-reminder"
	msgs "github.com/artyomturkin/protoactor-go-reminder/proto"
)

type exampleActor struct {
	reminder.Mixin

	wg *sync.WaitGroup
}

type ExampleMessage struct{}

func (r *exampleActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *ExampleMessage:
		r.RemindMe("hello", 1*time.Millisecond, false)
	case *msgs.Remind:
		if msg.Name == "hello" {
			fmt.Printf("Received reminder %s\n", msg.Name)
			r.wg.Done()
		}
	}
}

func exampleActorProducer(wg *sync.WaitGroup) func() actor.Actor {
	return func() actor.Actor {
		return &exampleActor{wg: wg}
	}
}

func ExampleMiddleware() {
	remProps := actor.FromProducer(reminder.Producer(10 * time.Millisecond)).
		WithMiddleware(persistence.Using(newProvider(5)))
	rem, err := actor.SpawnNamed(remProps, "reminder")
	if err != nil {
		fmt.Printf("failed to spawn reminder: %v", err)
		os.Exit(1)
	}

	wg := &sync.WaitGroup{}
	recProps := actor.FromProducer(exampleActorProducer(wg)).
		WithMiddleware(reminder.Middleware(rem))
	rec, err := actor.SpawnNamed(recProps, "receiver")
	if err != nil {
		fmt.Printf("failed to spawn receiver: %v", err)
		os.Exit(1)
	}

	rems := 2
	wg.Add(rems)
	for i := 0; i < rems; i++ {
		rec.Tell(&ExampleMessage{})
	}
	wg.Wait()

	rem.GracefulPoison()
	rec.GracefulPoison()

	// Output:
	// Received reminder hello
	// Received reminder hello
}

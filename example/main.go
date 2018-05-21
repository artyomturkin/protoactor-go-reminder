package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/persistence"
	"github.com/gogo/protobuf/types"

	reminder "github.com/artyomturkin/protoactor-go-reminder"

	msgs "github.com/artyomturkin/protoactor-go-reminder/proto"
)

// *** Example Actor ***
type receiver struct {
	wg *sync.WaitGroup
}

func (r *receiver) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *msgs.Remind:
		fmt.Printf("Got a reminder %s\n", msg.Name)
		r.wg.Done()
	}
}

func producer(wg *sync.WaitGroup) func() actor.Actor {
	return func() actor.Actor {
		return &receiver{wg: wg}
	}
}

// *** Example Persistence Store ***
type provider struct {
	providerState persistence.ProviderState
}

func newProvider(snapshotInterval int) *provider {
	return &provider{
		providerState: persistence.NewInMemoryProvider(snapshotInterval),
	}
}

func (p *provider) GetState() persistence.ProviderState {
	return p.providerState
}

// *** Example ***
func main() {
	//Create reminder actor with 10 ms trigger window and InMemory Persistence store
	remProps := actor.FromProducer(reminder.Producer(10 * time.Millisecond)).
		WithMiddleware(persistence.Using(newProvider(5)))
	rem, err := actor.SpawnNamed(remProps, "reminder")
	if err != nil {
		fmt.Printf("failed to spawn reminder: %v", err)
	}

	//Create and actor that will receive reminder
	wg := &sync.WaitGroup{}
	recProps := actor.FromProducer(producer(wg))
	rec, err := actor.SpawnNamed(recProps, "receiver")
	if err != nil {
		fmt.Printf("failed to spawn receiver: %v", err)
	}

	//Create timestamp of when reminder should trigger
	ti, _ := types.TimestampProto(time.Now().Add(1 * time.Millisecond))

	//Tell reminder actor to handle reminder. Use waitgroup to sync with receiver actor
	wg.Add(1)
	rem.Tell(&msgs.Reminder{
		Receiver: rec,
		At:       ti,
		Name:     "hello",
	})
	wg.Wait()

	//Shutdown all
	rec.GracefulPoison()
	rem.GracefulPoison()

	// Output:
	// Got a reminder hello
}

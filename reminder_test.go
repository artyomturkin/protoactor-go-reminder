package reminder_test

import (
	"sync"
	"testing"
	"time"

	"github.com/AsynkronIT/protoactor-go/persistence"

	"github.com/AsynkronIT/protoactor-go/actor"
	reminder "github.com/artyomturkin/protoactor-go-reminder"
	msgs "github.com/artyomturkin/protoactor-go-reminder/proto"
	protoTypes "github.com/gogo/protobuf/types"
)

type receiver struct {
	wg *sync.WaitGroup
}

func (r *receiver) Receive(ctx actor.Context) {
	switch ctx.Message().(type) {
	case *msgs.Reminded:
		r.wg.Done()
	}
}

func producer(wg *sync.WaitGroup) func() actor.Actor {
	return func() actor.Actor {
		return &receiver{wg: wg}
	}
}

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

//TestReminder test will fail with `all goroutines asleep` if reminder deos not work properly
func TestReminder(t *testing.T) {
	remProps := actor.FromProducer(reminder.Producer(10 * time.Millisecond)).
		WithMiddleware(persistence.Using(newProvider(5)))
	wg := &sync.WaitGroup{}
	recProps := actor.FromProducer(producer(wg))

	rem, err := actor.SpawnNamed(remProps, "reminder")
	if err != nil {
		t.Fatalf("failed to spawn reminder: %v", err)
	}
	rec, err := actor.SpawnNamed(recProps, "receiver")
	if err != nil {
		t.Fatalf("failed to spawn receiver: %v", err)
	}

	ti, _ := protoTypes.TimestampProto(time.Now().Add(1 * time.Millisecond))
	msg, _ := protoTypes.MarshalAny(&msgs.Reminded{At: ti})

	rems := 1
	wg.Add(rems)
	for i := 0; i < rems; i++ {
		rem.Tell(&msgs.Remind{
			Receiver: rec,
			At:       ti,
			Message:  msg,
		})
	}
	wg.Wait()

	rems = 10
	wg.Add(rems)
	for i := 0; i < rems; i++ {
		time.Sleep(1 * time.Millisecond)
		rem.Tell(&msgs.Remind{
			Receiver: rec,
			At:       ti,
			Message:  msg,
		})
	}
	wg.Wait()
}

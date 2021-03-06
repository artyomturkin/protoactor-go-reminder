package reminder_test

import (
	"sync"
	"testing"
	"time"

	"github.com/AsynkronIT/protoactor-go/persistence"

	"github.com/AsynkronIT/protoactor-go/actor"
	reminder "github.com/artyomturkin/protoactor-go-reminder"
	"github.com/gogo/protobuf/types"
)

type receiver struct {
	wg *sync.WaitGroup
}

var (
	mu      sync.Mutex = sync.Mutex{}
	counter int        = 0
)

func (r *receiver) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *reminder.Remind:
		if msg.Name == "hello" {
			mu.Lock()
			counter = counter + 1
			mu.Unlock()
			r.wg.Done()
		}
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

	ti, _ := types.TimestampProto(time.Now().Add(1 * time.Millisecond))

	rems := 1
	wg.Add(rems)
	for i := 0; i < rems; i++ {
		rem.Tell(&reminder.Reminder{
			Receiver: rec,
			At:       ti,
			Name:     "hello",
		})
	}
	wg.Wait()

	rems = 10
	wg.Add(rems)
	for i := 0; i < rems; i++ {
		time.Sleep(1 * time.Millisecond)
		rem.Tell(&reminder.Reminder{
			Receiver: rec,
			At:       ti,
			Name:     "hello",
		})
	}
	wg.Wait()

	//Test collation
	rems = 2
	counter = 0
	wg.Add(1)
	for i := 0; i < rems; i++ {
		time.Sleep(1 * time.Millisecond)
		ti, _ = types.TimestampProto(time.Now().Add(1 * time.Millisecond))
		rem.Tell(&reminder.Reminder{
			Receiver: rec,
			At:       ti,
			Name:     "hello",
			Collate:  true,
		})
	}
	wg.Wait()
	time.Sleep(50 * time.Millisecond)

	if counter != 1 {
		t.Errorf("collation failed! expected 1 got %d", counter)
	}

	rem.GracefulPoison()
	rec.GracefulPoison()
}

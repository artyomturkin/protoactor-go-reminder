package reminder

import (
	"context"
	"time"

	protoActor "github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/persistence"
	msgs "github.com/artyomturkin/protoactor-go-reminder/proto"
	protoTypes "github.com/gogo/protobuf/types"
)

type actor struct {
	persistence.Mixin
	self *protoActor.PID

	reminded time.Time
	reminds  []*msgs.Reminder

	triggersAt  time.Time
	cancelDelay func()

	window time.Duration
}

//Producer Reminder ActorProducer
func Producer(window time.Duration) func() protoActor.Actor {
	return func() protoActor.Actor { return &actor{window: window} }
}

var _ protoActor.Actor = (*actor)(nil)

func (a *actor) Receive(ctx protoActor.Context) {
	switch msg := ctx.Message().(type) {
	case *protoActor.Started:
		a.reminds = []*msgs.Reminder{}
		a.self = ctx.Self()
	case *msgs.Reminder:
		first := a.insertRemind(msg)

		if !a.Recovering() {
			a.PersistReceive(msg)

			if first {
				t, _ := protoTypes.TimestampFromProto(msg.At)
				delay := t.Add(a.window).Sub(time.Now())
				a.setDelay(delay)
			}
		}
	case *msgs.Reminded:
		rems := a.removeStale(msg)
		a.cancelDelay = nil

		if !a.Recovering() {
			for _, rem := range rems {
				m := &msgs.Remind{
					Name: rem.Name,
				}
				rem.Receiver.Tell(m)
			}
			a.PersistReceive(msg)
		}
	case *msgs.Snapshot:
		a.reminds = msg.Reminds
		t, _ := protoTypes.TimestampFromProto(msg.At)
		a.reminded = t
	case *persistence.ReplayComplete:
		if len(a.reminds) > 0 {
			t, _ := protoTypes.TimestampFromProto(a.reminds[0].At)
			delay := t.Add(a.window).Sub(time.Now())
			a.setDelay(delay)
		}
	case *persistence.RequestSnapshot:
		t, _ := protoTypes.TimestampProto(a.reminded)
		snap := &msgs.Snapshot{
			Reminds: a.reminds,
			At:      t,
		}
		a.PersistSnapshot(snap)
	}
}

func (a *actor) removeStale(msg *msgs.Reminded) []*msgs.Reminder {
	t, _ := protoTypes.TimestampFromProto(msg.At)
	stale := []*msgs.Reminder{}
	found := false
	for i, rem := range a.reminds {
		curT, _ := protoTypes.TimestampFromProto(rem.At)
		if t.Before(curT) {
			found = true
			stale = a.reminds[:i+1]
			a.reminds = a.reminds[i+1:]
			break
		}
	}
	if !found {
		stale = a.reminds
		a.reminds = []*msgs.Reminder{}
	}

	return stale
}

func (a *actor) insertRemind(msg *msgs.Reminder) (first bool) {
	inserted := false
	first = false
	for i, rem := range a.reminds {
		msgT, _ := protoTypes.TimestampFromProto(msg.At)
		curT, _ := protoTypes.TimestampFromProto(rem.At)
		if msgT.Before(curT) {
			newRems := append(a.reminds[:i], msg)
			a.reminds = append(newRems, a.reminds[i:]...)
			inserted = true
			if i == 0 {
				first = true
			}
			break
		}
	}
	if !inserted {
		a.reminds = append(a.reminds, msg)
	}

	return first || len(a.reminds) == 1
}

func (a *actor) setDelay(t time.Duration) {
	triggersAt := time.Now().Add(t)
	if a.cancelDelay == nil || (triggersAt.Before(a.triggersAt)) {
		if a.cancelDelay != nil {
			a.cancelDelay()
		}
		a.triggersAt = triggersAt

		ctx, cancel := context.WithCancel(context.Background())
		a.cancelDelay = cancel

		go func() {
			select {
			case <-ctx.Done():
				break
			case <-time.After(t):
				ts, _ := protoTypes.TimestampProto(triggersAt)
				a.self.Tell(&msgs.Reminded{
					At: ts,
				})
				break
			}
		}()
	}
}

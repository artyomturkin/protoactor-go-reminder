package reminder

import (
	"context"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/persistence"
	"github.com/gogo/protobuf/types"
)

type reminderActor struct {
	persistence.Mixin
	self *actor.PID

	reminded time.Time
	reminds  []*Reminder

	triggersAt  time.Time
	cancelDelay func()

	window time.Duration
}

//Producer Reminder ActorProducer
func Producer(window time.Duration) func() actor.Actor {
	return func() actor.Actor { return &reminderActor{window: window} }
}

var _ actor.Actor = (*reminderActor)(nil)

func (a *reminderActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		a.reminds = []*Reminder{}
		a.self = ctx.Self()
	case *Reminder:
		a.collate(msg)

		first := a.insertRemind(msg)

		if !a.Recovering() {
			a.PersistReceive(msg)

			if first {
				t, _ := types.TimestampFromProto(msg.At)
				delay := t.Add(a.window).Sub(time.Now())
				a.setDelay(delay)
			}
		}
	case *Reminded:
		rems := a.removeStale(msg)
		a.cancelDelay = nil

		if !a.Recovering() {
			for _, rem := range rems {
				m := &Remind{
					Name: rem.Name,
				}
				rem.Receiver.Tell(m)
			}
			a.PersistReceive(msg)
		}
	case *Snapshot:
		a.reminds = msg.Reminds
		t, _ := types.TimestampFromProto(msg.At)
		a.reminded = t
	case *persistence.ReplayComplete:
		if len(a.reminds) > 0 {
			t, _ := types.TimestampFromProto(a.reminds[0].At)
			delay := t.Add(a.window).Sub(time.Now())
			a.setDelay(delay)
		}
	case *persistence.RequestSnapshot:
		t, _ := types.TimestampProto(a.reminded)
		snap := &Snapshot{
			Reminds: a.reminds,
			At:      t,
		}
		a.PersistSnapshot(snap)
	}
}

func (a *reminderActor) removeStale(msg *Reminded) []*Reminder {
	t, _ := types.TimestampFromProto(msg.At)
	stale := []*Reminder{}
	found := false
	for i, rem := range a.reminds {
		curT, _ := types.TimestampFromProto(rem.At)
		if t.Before(curT) {
			found = true
			stale = a.reminds[:i+1]
			a.reminds = a.reminds[i+1:]
			break
		}
	}
	if !found {
		stale = a.reminds
		a.reminds = []*Reminder{}
	}

	return stale
}

func (a *reminderActor) insertRemind(msg *Reminder) (first bool) {
	inserted := false
	first = false
	for i, rem := range a.reminds {
		msgT, _ := types.TimestampFromProto(msg.At)
		curT, _ := types.TimestampFromProto(rem.At)
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

func (a *reminderActor) setDelay(t time.Duration) {
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
				ts, _ := types.TimestampProto(triggersAt)
				a.self.Tell(&Reminded{
					At: ts,
				})
				break
			}
		}()
	}
}

func (a *reminderActor) collate(msg *Reminder) {
	if msg.Collate {
		for i, r := range a.reminds {
			if r.Name == msg.Name && r.Receiver.Equal(msg.Receiver) && r.Collate {
				a.reminds = append(a.reminds[:i], a.reminds[i+1:]...)
			}
		}
	}
}

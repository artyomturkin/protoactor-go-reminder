package reminder

import (
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/gogo/protobuf/types"
)

type remindable interface {
	init(self, reminder *actor.PID)
	RemindMe(string, time.Duration, bool)
}

//Mixin provides reminder service to actor
type Mixin struct {
	remActor *actor.PID
	self     *actor.PID
}

var _ remindable = (*Mixin)(nil)

func (m *Mixin) init(self, reminder *actor.PID) {
	m.remActor = reminder
	m.self = self
}

//RemindMe set reminder
func (m *Mixin) RemindMe(name string, in time.Duration, collate bool) {
	ti, _ := types.TimestampProto(time.Now().Add(in))

	m.remActor.Tell(&Reminder{
		Receiver: m.self,
		At:       ti,
		Name:     name,
		Collate:  collate,
	})
}

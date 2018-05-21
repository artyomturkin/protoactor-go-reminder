package reminder

import (
	"time"

	protoActor "github.com/AsynkronIT/protoactor-go/actor"
	msgs "github.com/artyomturkin/protoactor-go-reminder/proto"
	protoTypes "github.com/gogo/protobuf/types"
)

type remindable interface {
	init(self, reminder *protoActor.PID)
	RemindMe(string, time.Time, bool)
}

//Mixin provides reminder service to actor
type Mixin struct {
	remActor *protoActor.PID
	self     *protoActor.PID
}

var _ remindable = (*Mixin)(nil)

func (m *Mixin) init(self, reminder *protoActor.PID) {
	m.remActor = reminder
	m.self = self
}

//RemindMe set reminder
func (m *Mixin) RemindMe(name string, at time.Time, collate bool) {
	ti, _ := protoTypes.TimestampProto(at)

	m.remActor.Tell(&msgs.Reminder{
		Receiver: m.self,
		At:       ti,
		Name:     name,
		Collate:  collate,
	})
}

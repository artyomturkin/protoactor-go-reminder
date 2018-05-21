package reminder

import (
	"log"
	"reflect"

	protoActor "github.com/AsynkronIT/protoactor-go/actor"
)

//Reminder middleware
func Reminder(reminder *protoActor.PID) func(next protoActor.ActorFunc) protoActor.ActorFunc {
	return func(next protoActor.ActorFunc) protoActor.ActorFunc {
		return func(ctx protoActor.Context) {
			switch ctx.Message().(type) {
			case *protoActor.Started:
				if p, ok := ctx.Actor().(remindable); ok {
					p.init(ctx.Self(), reminder)
				} else {
					log.Fatalf("Actor type %v is not remindable", reflect.TypeOf(ctx.Actor()))
				}
				next(ctx)
			default:
				next(ctx)
			}
		}
	}
}

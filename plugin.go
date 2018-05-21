package reminder

import (
	"log"
	"reflect"

	"github.com/AsynkronIT/protoactor-go/actor"
)

//Middleware middleware
func Middleware(reminder *actor.PID) func(next actor.ActorFunc) actor.ActorFunc {
	return func(next actor.ActorFunc) actor.ActorFunc {
		return func(ctx actor.Context) {
			switch ctx.Message().(type) {
			case *actor.Started:
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

# Proto Actor [Go] - Reminder

[![Go Report Card](https://goreportcard.com/badge/github.com/artyomturkin/protoactor-go-reminder)](https://goreportcard.com/report/github.com/artyomturkin/protoactor-go-reminder)

Go package that provides a mechanism for actors to receive time delayed reminders.

Idea based on [Microsoft Orleans Reminders](https://dotnet.github.io/orleans/Documentation/Core-Features/Timers-and-Reminders.html)

## Get started

Install package:

```
go get github.com/artyomturkin/protoactor-go-reminder
```

Create a reminder actor. **It requires persistence plugin to work.**

```go
func main() {
	//Create reminder actor with 10 ms trigger window and InMemory Persistence store
	remProps := actor.FromProducer(reminder.Producer(10 * time.Millisecond)).
		WithMiddleware(persistence.Using(newProvider(5)))
	rem, err := actor.SpawnNamed(remProps, "reminder")
	if err != nil {
		fmt.Printf("failed to spawn reminder: %v", err)
		os.Exit(1)
	}

	.....
}
```

Define actor, that can setup and receive reminders
```go
type exampleActor struct {
	reminder.Mixin
}

type ExampleMessage struct{}

func (r *exampleActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *ExampleMessage:
		r.RemindMe("hello", 1*time.Millisecond, false)
	case *msgs.Remind:
		if msg.Name == "hello" {
			fmt.Printf("Received reminder %s\n", msg.Name)
		}
	}
}
```

Create example actor instance
```go
func main() {
	.....
	recProps := actor.FromProducer(exampleActorProducer).
		WithMiddleware(reminder.Reminder(rem))
	rec, err := actor.SpawnNamed(recProps, "receiver")
	if err != nil {
		fmt.Printf("failed to spawn receiver: %v", err)
		os.Exit(1)
	}
	.....
}
```

`Tell` example actor to setup a remind for itself
```go
func main() {
	.....

	rec.Tell(&ExampleMessage{})
	
	.....
}
```
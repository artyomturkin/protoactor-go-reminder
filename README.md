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
	}

	.....
}
```

Use `import "github.com/gogo/protobuf/types"` package to create reminder time
```go
func main() {
	.....

	//Create timestamp of when reminder should trigger
	ti, _ := types.TimestampProto(time.Now().Add(1 * time.Millisecond))

	.....
}
```

`Tell` reminder actor to handle a remind
```go
func main() {
	.....

	rem.Tell(&msgs.Reminder{
		Receiver: rec,
		At:	     ti,
		Name:    "hello",
		Collate: false,
	})
	
	.....
}
```
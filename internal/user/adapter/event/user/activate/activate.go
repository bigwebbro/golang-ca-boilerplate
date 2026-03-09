package activate

import (
	"example.com/golang-boilerplate/user/application/dispatcher"
	domainEvent "example.com/pkg/event"
	"fmt"
	"github.com/gookit/event"
)

type EventHandler struct {
	domainEvent domainEvent.Emitter
}

func New() *EventHandler {
	return &EventHandler{}
}

func (h EventHandler) Emit(e domainEvent.Emitter) {
	//TODO implement me
	panic("implement me")
}

func (h EventHandler) EventHandler() dispatcher.Emitter {
	event.On("user-activate", event.ListenerFunc(func(e event.Event) error {
		fmt.Printf("handle event: %s\n", e.Name())
		return nil
	}), event.High)
	panic("123")
}

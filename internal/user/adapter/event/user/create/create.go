package create

import (
	"encoding/json"
	"example.com/golang-boilerplate/user/adapter/storage/psql"
	"example.com/golang-boilerplate/user/application/dispatcher"
	"example.com/golang-boilerplate/user/domain/user/event/create"
	app_event "example.com/pkg/event"
	"fmt"
	"github.com/gookit/event"
)

type EventHandler struct {
	storage psql.Adapter
}

func New(storage psql.Adapter) *EventHandler {
	return &EventHandler{storage: storage}
}

func (h EventHandler) Emit(e app_event.Emitter) {
	event.MustFire(e.Type(), map[string]interface{}{"event": e.Emit()})
}

func (h EventHandler) EventHandler() dispatcher.Emitter {
	event.On(create.Name, event.ListenerFunc(func(e event.Event) error {
		var create = &create.Create{}
		b, _ := e.Get("event").([]byte)
		json.Unmarshal(b, create)
		h.storage.CreateUser(*create)
		fmt.Printf("handle event: %s\n", e.Name())

		return nil
	}), event.High)

	return h
}

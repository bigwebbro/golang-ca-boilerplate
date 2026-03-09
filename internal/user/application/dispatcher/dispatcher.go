package dispatcher

import (
	"example.com/pkg/event"
)

//in-adapter -> query -> provider (load) -> adapter
//in-adapter -> usecase -> provider (load) -> domain -> event -> dispatcher -> ...adapter

type EventDispatcher interface {
	AddEmitter(e string, h Emitter)
	Dispatch(es ...event.Emitter)
}

type Emitter interface {
	Emit(e event.Emitter)
}

type Dispatcher struct {
	routing map[string]Emitter
}

func New() *Dispatcher {
	return &Dispatcher{routing: map[string]Emitter{}}
}

func (d Dispatcher) AddEmitter(e string, h Emitter) {
	d.routing[e] = h
}

func (d Dispatcher) Dispatch(es ...event.Emitter) {
	for _, e := range es {
		emitter := d.routing[e.Type()]
		emitter.Emit(e)
	}
}

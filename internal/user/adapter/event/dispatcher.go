package event

import (
	"context"

	"example.com/golang-boilerplate/user/application/dispatcher"
)

type ContextAware interface {
	SetContext(ctx context.Context)
}

type RegisterHandler interface {
	EventHandler() dispatcher.Emitter
}

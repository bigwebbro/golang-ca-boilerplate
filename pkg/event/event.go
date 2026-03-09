package event

import (
	"github.com/google/uuid"
	"time"
)

type Event struct {
	ID        uuid.UUID
	EntityID  uuid.UUID
	CreatedAt time.Time
}

type Emitter interface {
	Emit() []byte
	Type() string
	GetEvent() Event
}

package group

import (
	"example.com/golang-boilerplate/user/domain/group/type/name"
	"github.com/google/uuid"
)

type Group struct {
	ID   uuid.UUID
	name name.Name
}

package query

import (
	"context"
	"example.com/golang-boilerplate/user/domain/group"
	"example.com/golang-boilerplate/user/domain/user"
	"github.com/google/uuid"
)

type User interface {
	User(ctx context.Context, id ...uuid.UUID) []user.User
}

type Group interface {
	Group(ctx context.Context, id ...uuid.UUID) []group.Group
}

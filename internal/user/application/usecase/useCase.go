package usecase

import (
	"context"

	"example.com/golang-boilerplate/user/domain/group"
	"example.com/golang-boilerplate/user/domain/user"
	"example.com/golang-boilerplate/user/domain/user/type/age"
	"example.com/golang-boilerplate/user/domain/user/type/email"
	"example.com/golang-boilerplate/user/domain/user/type/name"
	"example.com/golang-boilerplate/user/domain/user/type/patronymic"
	"example.com/golang-boilerplate/user/domain/user/type/surname"
	"github.com/google/uuid"
)

type Group interface {
	Create() group.Group
}

type User interface {
	Create(c UserCreateCommand) (*user.User, error)
	Activate(ctx context.Context, uuid uuid.UUID) user.User
	DeActivate(ctx context.Context, uuid uuid.UUID) user.User
	ChangeGroup(ctx context.Context, c UserChangeGroupCommand) group.Group
	RulesConfirmation(ctx context.Context, user user.User)
	MailingConsent(ctx context.Context, user user.User)
}

type UserCreateCommand struct {
	Name       name.Name
	Patronymic patronymic.Patronymic
	Surname    surname.Surname
	Age        age.Age
	Email      email.Email
	Group      group.Group
}

type UserChangeGroupCommand struct {
	userId  uuid.UUID
	groupId uuid.UUID
}

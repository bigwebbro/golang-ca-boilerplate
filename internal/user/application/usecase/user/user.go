package user

import (
	"context"
	"example.com/golang-boilerplate/user/application/dispatcher"
	"example.com/golang-boilerplate/user/application/provider"
	"example.com/golang-boilerplate/user/application/usecase"
	"example.com/golang-boilerplate/user/domain/group"
	"example.com/golang-boilerplate/user/domain/user"
	"example.com/pkg/errors"
	"github.com/google/uuid"
)

type UseCase struct {
	groupProvider provider.Group
	userProvider  provider.User
	dispatcher    dispatcher.Dispatcher
}

func New(groupProvider provider.Group, userProvider provider.User, dispatcher *dispatcher.Dispatcher) *UseCase {
	return &UseCase{groupProvider: groupProvider, userProvider: userProvider, dispatcher: *dispatcher}
}

func (uc UseCase) Create(c usecase.UserCreateCommand) (*user.User, error) {
	u, _ := uc.userProvider.UserByEmail(c.Email.Value())

	if u != nil {
		return nil, errors.Newf("User with email %s is not uniq", c.Email)
	}

	u = user.New(c.Name, c.Patronymic, c.Surname, c.Age, c.Email)
	uc.dispatcher.Dispatch(u.Events()...)

	return u, nil
}

func (u UseCase) Activate(ctx context.Context, uuid uuid.UUID) user.User {
	//TODO implement me
	panic("implement me")
}

func (u UseCase) DeActivate(ctx context.Context, uuid uuid.UUID) user.User {
	//TODO implement me
	panic("implement me")
}

func (u UseCase) ChangeGroup(ctx context.Context, c usecase.UserChangeGroupCommand) group.Group {
	//TODO implement me
	panic("implement me")
}

func (u UseCase) RulesConfirmation(ctx context.Context, user user.User) {
	//TODO implement me
	panic("implement me")
}

func (u UseCase) MailingConsent(ctx context.Context, user user.User) {
	//TODO implement me
	panic("implement me")
}

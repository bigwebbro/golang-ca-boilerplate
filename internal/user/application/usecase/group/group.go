package group

import (
	"example.com/golang-boilerplate/user/application/dispatcher"
	"example.com/golang-boilerplate/user/application/provider"
	"example.com/golang-boilerplate/user/domain/group"
)

type UseCase struct {
	groupProvider provider.Group
	dispatcher    *dispatcher.Dispatcher
}

func New(groupProvider provider.Group, dispatcher *dispatcher.Dispatcher) *UseCase {
	return &UseCase{groupProvider: groupProvider, dispatcher: dispatcher}
}

func (u UseCase) Create() group.Group {
	panic("implement me")
}

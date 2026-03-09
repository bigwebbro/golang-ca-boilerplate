package psql

import (
	"example.com/golang-boilerplate/user/domain/group"
	group_create "example.com/golang-boilerplate/user/domain/group/event/create"
)

type Group struct {
}

func (a Adapter) Group() group.Group {
	//TODO implement me
	panic("implement me")
}

func (a Adapter) CreateGroup(e group_create.Create) {
	//TODO implement me
	panic("implement me")
}

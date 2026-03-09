package create

import (
	"example.com/golang-boilerplate/user/domain/group"
	"example.com/pkg/event"
)

type Create struct {
	event.Event
	*group.Group
}

func (c Create) Emit() {

}

package change_group

import (
	"example.com/golang-boilerplate/user/domain/user"
	"example.com/pkg/event"
)

type ChangeGroup struct {
	event.Event
	*user.User
}

func (c ChangeGroup) Emit() {

}

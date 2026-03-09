package deactivate

import (
	"example.com/golang-boilerplate/user/domain/user"
	"example.com/pkg/event"
)

type DeActivate struct {
	event.Event
	*user.User
}

func (c DeActivate) Emit() {

}

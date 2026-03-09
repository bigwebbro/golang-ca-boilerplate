package activate

import (
	"example.com/golang-boilerplate/user/domain/user"
	"example.com/pkg/event"
)

const name = "user:activate"

type Activate struct {
	event.Event
	*user.User
}

func (c Activate) Emit() {

}

func (c Activate) Type() string {
	return name
}

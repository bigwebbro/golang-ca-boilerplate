package rules_confirmation

import (
	"example.com/golang-boilerplate/user/domain/user"
	"example.com/pkg/event"
)

type RulesConfirmation struct {
	event.Event
	*user.User
}

func (c RulesConfirmation) Emit() {

}

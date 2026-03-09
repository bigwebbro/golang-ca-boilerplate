package mailing_consent

import (
	"example.com/golang-boilerplate/user/domain/user"
	"example.com/pkg/event"
)

type MailingConsent struct {
	event.Event
	*user.User
}

func (c MailingConsent) Emit() {

}

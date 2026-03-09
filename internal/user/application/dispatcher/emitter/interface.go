package emitter

// out interfaces
import (
	group_create "example.com/golang-boilerplate/user/domain/group/event/create"
	"example.com/golang-boilerplate/user/domain/user/event/activate"
	"example.com/golang-boilerplate/user/domain/user/event/change_group"
	"example.com/golang-boilerplate/user/domain/user/event/create"
	"example.com/golang-boilerplate/user/domain/user/event/deactivate"
	"example.com/golang-boilerplate/user/domain/user/event/mailing_consent"
	"example.com/golang-boilerplate/user/domain/user/event/rules_confirmation"
)

type User interface {
	CreateUser(e create.Create)
	Activate(e activate.Activate)
	DeActivate(e deactivate.DeActivate)
	ChangeGroup(e change_group.ChangeGroup)
	MailingConsent(e mailing_consent.MailingConsent)
	RulesConfirmation(e rules_confirmation.RulesConfirmation)
}

type Group interface {
	CreateGroup(e group_create.Create)
}

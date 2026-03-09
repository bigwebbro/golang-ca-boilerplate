package psql

import (
	"example.com/golang-boilerplate/user/domain/user"
	"example.com/golang-boilerplate/user/domain/user/event/activate"
	"example.com/golang-boilerplate/user/domain/user/event/change_group"
	"example.com/golang-boilerplate/user/domain/user/event/create"
	"example.com/golang-boilerplate/user/domain/user/event/deactivate"
	"example.com/golang-boilerplate/user/domain/user/event/mailing_consent"
	"example.com/golang-boilerplate/user/domain/user/event/rules_confirmation"
	"example.com/golang-boilerplate/user/domain/user/type/email"
	"example.com/golang-boilerplate/user/domain/user/type/name"
	"example.com/golang-boilerplate/user/domain/user/type/patronymic"
	"example.com/golang-boilerplate/user/domain/user/type/surname"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type User struct {
	ID         uuid.UUID `gorm:"primarykey;type:uuid"`
	Name       string
	Patronymic string
	Surname    string
	Email      string `gorm:"uniqueIndex:idx_email"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (a Adapter) CreateUser(e create.Create) {
	var u = &User{
		ID:         e.GetEvent().EntityID,
		Name:       e.Name,
		Patronymic: e.Patronymic,
		Surname:    e.Surname,
		Email:      e.Email,
		CreatedAt:  e.GetEvent().CreatedAt,
	}

	a.db.Save(u)
}

func (a Adapter) Activate(e activate.Activate) {
	//TODO implement me
	panic("implement me")
}

func (a Adapter) DeActivate(e deactivate.DeActivate) {
	//TODO implement me
	panic("implement me")
}

func (a Adapter) ChangeGroup(e change_group.ChangeGroup) {
	//TODO implement me
	panic("implement me")
}

func (a Adapter) MailingConsent(e mailing_consent.MailingConsent) {
	//TODO implement me
	panic("implement me")
}

func (a Adapter) RulesConfirmation(e rules_confirmation.RulesConfirmation) {
	//TODO implement me
	panic("implement me")
}

// UserByEmail Provider Interface
func (a Adapter) UserByEmail(email string) (*user.User, error) {
	var u = &User{}
	a.db.Where(&User{Email: email}).First(u)

	if len(u.Email) > 0 {
		return u.toDomain(), nil
	}

	return nil, errors.Errorf("Entity not found")
}

func (u User) toDomain() *user.User {
	n, _ := name.New(u.Name)
	p, _ := patronymic.New(u.Patronymic)
	s, _ := surname.New(u.Surname)
	e, _ := email.New(u.Email)

	return user.NewWithID(u.ID, n, p, s, e, u.CreatedAt)
}

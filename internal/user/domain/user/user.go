package user

import (
	"example.com/golang-boilerplate/user/domain/group"
	"example.com/golang-boilerplate/user/domain/user/event/create"
	"example.com/golang-boilerplate/user/domain/user/type/age"
	"example.com/golang-boilerplate/user/domain/user/type/email"
	"example.com/golang-boilerplate/user/domain/user/type/name"
	"example.com/golang-boilerplate/user/domain/user/type/patronymic"
	"example.com/golang-boilerplate/user/domain/user/type/surname"
	"example.com/pkg/event"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID         uuid.UUID
	name       name.Name
	patronymic patronymic.Patronymic
	surname    surname.Surname
	age        age.Age
	email      email.Email
	group      group.Group
	events     []event.Emitter
	createdAt  time.Time
	updatedAt  time.Time
}

func New(name name.Name, patronymic patronymic.Patronymic, surname surname.Surname, age age.Age, email email.Email) *User {
	var timeNow = time.Now().UTC()
	u := User{
		ID:         uuid.New(),
		name:       name,
		patronymic: patronymic,
		surname:    surname,
		age:        age,
		email:      email,
		events:     []event.Emitter{},
		createdAt:  timeNow,
	}

	e := create.New(u.ID, u.name, u.patronymic, u.surname, u.age, u.email)

	u.events = append(u.events, e)

	return &u
}

func NewWithID(ID uuid.UUID, name name.Name, patronymic patronymic.Patronymic, surname surname.Surname, email email.Email, createdAt time.Time) *User {
	return &User{
		ID:         ID,
		name:       name,
		patronymic: patronymic,
		surname:    surname,
		email:      email,
		events:     []event.Emitter{},
		createdAt:  createdAt,
		updatedAt:  time.Now(),
	}
}

func (u User) Name() name.Name {
	return u.name
}

func (u User) Patronymic() patronymic.Patronymic {
	return u.patronymic
}

func (u User) Surname() surname.Surname {
	return u.surname
}

func (u User) Events() []event.Emitter {
	return u.events
}

func (u User) Email() email.Email {
	return u.email
}

func (u User) CreatedAt() time.Time {
	return u.createdAt
}

func (u User) UpdatedAt() time.Time {
	return u.updatedAt
}

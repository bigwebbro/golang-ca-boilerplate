package create

import (
	"encoding/json"
	"example.com/golang-boilerplate/user/domain/user/type/age"
	"example.com/golang-boilerplate/user/domain/user/type/email"
	"example.com/golang-boilerplate/user/domain/user/type/name"
	"example.com/golang-boilerplate/user/domain/user/type/patronymic"
	"example.com/golang-boilerplate/user/domain/user/type/surname"
	"example.com/pkg/event"
	"github.com/google/uuid"
)

const Name = "user-create"

type Create struct {
	event.Event
	Name       string
	Patronymic string
	Surname    string
	Age        uint16
	Email      string
}

func New(ID uuid.UUID, name name.Name, patronymic patronymic.Patronymic, surname surname.Surname, age age.Age, email email.Email) *Create {
	return &Create{
		Event: event.Event{
			ID:       uuid.New(),
			EntityID: ID,
		},
		Name:       name.Value(),
		Patronymic: patronymic.Value(),
		Surname:    surname.Value(),
		Age:        age.Value(),
		Email:      email.Value(),
	}
}

func (c Create) Emit() []byte {
	b, _ := json.Marshal(c)

	return b
}

func (c Create) Type() string {
	return Name
}

func (c Create) GetEvent() event.Event {
	return c.Event
}

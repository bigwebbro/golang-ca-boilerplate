package api

import (
	"example.com/golang-boilerplate/user/application/usecase"
	"example.com/golang-boilerplate/user/domain/user/type/age"
	"example.com/golang-boilerplate/user/domain/user/type/email"
	"example.com/golang-boilerplate/user/domain/user/type/name"
	"example.com/golang-boilerplate/user/domain/user/type/patronymic"
	"example.com/golang-boilerplate/user/domain/user/type/surname"
)

type CreateUserRequest struct {
	Name       string `json:"name"`
	Patronymic string `json:"patronymic"`
	Surname    string `json:"surname"`
	Age        uint64 `json:"age"`
	Email      string `json:"email"`
}

type errsMsg []string

func (r CreateUserRequest) toCommand() (usecase.UserCreateCommand, errsMsg) {
	var errs errsMsg

	n, err := name.New(r.Name)
	if err != nil {
		errs = append(errs, err.Error())
	}

	p, err := patronymic.New(r.Patronymic)
	if err != nil {
		errs = append(errs, err.Error())
	}

	s, err := surname.New(r.Surname)
	if err != nil {
		errs = append(errs, err.Error())
	}

	a, err := age.New(r.Age)
	if err != nil {
		errs = append(errs, err.Error())
	}

	e, err := email.New(r.Email)
	if err != nil {
		errs = append(errs, err.Error())
	}

	return usecase.UserCreateCommand{
		Name:       n,
		Patronymic: p,
		Surname:    s,
		Age:        a,
		Email:      e,
	}, errs
}

package provider

import (
	"example.com/golang-boilerplate/user/domain/group"
	"example.com/golang-boilerplate/user/domain/user"
)

type User interface {
	UserByEmail(email string) (*user.User, error)
}

type Group interface {
	Group() group.Group
}

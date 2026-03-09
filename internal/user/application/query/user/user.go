package user

import (
	"example.com/golang-boilerplate/user/application/provider"
)

type Query struct {
	userProvider provider.Group
}

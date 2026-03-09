package api

import (
	"example.com/golang-boilerplate/user/application/usecase"
	"github.com/gofiber/fiber/v2"
)

type Adapter struct {
	ucUser  usecase.User
	ucGroup usecase.Group

	fiber   *fiber.App
	options options
}

func New(ucUser usecase.User, ucGroup usecase.Group, ops ...Option) *Adapter {
	var a = &Adapter{ucUser: ucUser, ucGroup: ucGroup}
	cfg := &options{}

	for _, opt := range ops {
		opt(cfg)
	}

	var api fiber.Router
	a.fiber, api = a.api(fiber.Config(*cfg))
	a.groupApi(a.userApi(api))

	return a
}

func (a Adapter) Run(addr string) error {
	return a.fiber.Listen(addr)
}

func (a Adapter) RunTLS(addr, cert, key string) error {
	return a.fiber.ListenTLS(addr, cert, key)
}

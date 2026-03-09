package api

import "github.com/gofiber/fiber/v2"

func (a Adapter) api(cfg fiber.Config) (*fiber.App, fiber.Router) {
	f := fiber.New(cfg)

	return f, f.Group("/api/v1")
}

func (a Adapter) userApi(r fiber.Router) fiber.Router {
	u := r.Group("/user")
	u.Post("", a.CreateUser)
	u.Put("/:id/activate", a.CreateUser)
	u.Put("/:id/deactivate", a.CreateUser)
	u.Put("/:id/:groupId", a.CreateUser)

	return u
}

func (a Adapter) groupApi(r fiber.Router) fiber.Router {
	g := r.Group("/group")
	g.Post("", a.CreateUser)

	return g
}

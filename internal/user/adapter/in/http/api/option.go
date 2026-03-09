package api

import "github.com/gofiber/fiber/v2"

type options fiber.Config

type Option func(*options)

func WithPrefork() Option {
	return func(o *options) {
		o.Prefork = true
	}
}

func WithCaseSensitive() Option {
	return func(o *options) {
		o.CaseSensitive = true
	}
}

func WithStrictRouting() Option {
	return func(o *options) {
		o.StrictRouting = true
	}
}

module example.com/golang-boilerplate/user/application

go 1.25.4

require (
	example.com/golang-boilerplate/user/domain v0.0.0-00010101000000-000000000000
	example.com/pkg v0.0.0-00010101000000-000000000000
	github.com/google/uuid v1.6.0
)

replace example.com/golang-boilerplate/user/domain => ../domain

replace example.com/pkg => ../../../pkg

module example.com/golang-boilerplate/user/adapter

go 1.25.4

require (
	example.com/golang-boilerplate/user/application v0.0.0-00010101000000-000000000000
	example.com/golang-boilerplate/user/domain v0.0.0-00010101000000-000000000000
	example.com/pkg v0.0.0-00010101000000-000000000000
	github.com/gofiber/fiber/v2 v2.52.12
	github.com/google/uuid v1.6.0
	github.com/gookit/event v1.2.0
	github.com/pkg/errors v0.9.1
	gorm.io/driver/postgres v1.6.0
	gorm.io/gorm v1.31.1
)

require (
	github.com/andybalholm/brotli v1.1.0 // indirect
	github.com/gookit/goutil v0.7.1 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.6.0 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.51.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/sync v0.11.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
)

replace example.com/golang-boilerplate/user/domain => ../../../internal/user/domain

replace example.com/golang-boilerplate/user/application =>  ../../../internal/user/application

replace example.com/golang-boilerplate/user/adapter =>  ../../../internal/user/adapter

replace example.com/pkg => ../../../pkg

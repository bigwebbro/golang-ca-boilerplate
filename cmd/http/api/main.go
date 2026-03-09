package main

import (
	"fmt"
	"log"

	"example.com/golang-boilerplate/http/api/user/config"
	createUserEH "example.com/golang-boilerplate/user/adapter/event/user/create"
	"example.com/golang-boilerplate/user/adapter/in/http/api"
	"example.com/golang-boilerplate/user/adapter/storage/psql"
	"example.com/golang-boilerplate/user/application/dispatcher"
	useCaseGroup "example.com/golang-boilerplate/user/application/usecase/group"
	useCaseUser "example.com/golang-boilerplate/user/application/usecase/user"
	"example.com/golang-boilerplate/user/domain/user/event/create"
)

func main() {
	viper := config.Viper()

	psqlAdapter := psql.New(
		viper.GetString("DB_HOST"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASS"),
		viper.GetString("DB_PORT"),
	)
	psqlAdapter.AutoMigration()

	d := dispatcher.New()
	d.AddEmitter(create.Name, createUserEH.New(psqlAdapter).EventHandler())

	var (
		ucUser  = useCaseUser.New(psqlAdapter, psqlAdapter, d)
		ucGroup = useCaseGroup.New(psqlAdapter, d)
		httpApi = api.New(ucUser, ucGroup, api.WithStrictRouting(), api.WithCaseSensitive())
	)

	if err := httpApi.Run(":" + viper.GetString("PORT")); err != nil {
		log.Panic(err)
	}

	fmt.Printf("service started successfully on http port: %s", viper.GetString("PORT"))
}

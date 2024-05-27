package main

import (
	user "auth/user/entities"
	"loki/thor"
)

func main() {
	db := thor.GetPostgresClient()
	db.AutoMigrate(
		user.User{},
	)
}

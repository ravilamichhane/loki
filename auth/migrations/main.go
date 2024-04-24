package main

import (
	user "auth/user/entities"
	"nest/thor"
)

func main() {
	db := thor.GetPostgresClient()
	db.AutoMigrate(
		user.User{},
	)
}

package main

import (
	"app/user/entities"
	"nest/thor"
)

func main() {
	db := thor.GetPostgresClient()
	db.AutoMigrate(
		entities.User{},
	)
}

package main

import (
	"auth/user"
	"context"
	"log"
	"nest/common"
	"nest/core"
	"nest/logger"
	"nest/thor"
	"os"
)

func main() {
	log.Print("Starting servers")

	thor.LoadEnv()

	logger := logger.New("builder-api", os.Stdout, logger.LevelInfo, thor.GetTraceID, logger.DiscordEvent(logger.DiscordEventConfig{
		WebhookURL: "https://discord.com/api/webhooks/1224246134516744317/15uoW7BG5Jv6Ai1sMvqEU2a8an76vwflPzlDDKEOtTro43OEXlT24Trxr_PCisod9JiF",
		Error:      true,
	}))

	db := thor.GetPostgresClient()

	userService := user.NewUserServiceDB(db)
	userController := user.NewUserController(userService)

	appModule := &common.Module{
		Controllers: []common.ControllerBase{
			userController,
		},
	}

	ctx := context.Background()

	c := core.NewNestFactory(thor.NewApp(ctx, thor.ThorConfig{
		Logger: logger,
	})).Create(appModule)

	

	c.Listen(":8080")

}

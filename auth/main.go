package main

import (
	"auth/iam/authentication"
	"auth/iam/authorisation"
	"auth/user"
	"context"
	"log"
	"nest/common"
	"nest/core"
	"nest/logger"
	"nest/thor"
	"net"
	"os"

	"google.golang.org/grpc"
)

type server struct {
	authorisation.UnimplementedAuthorisationServiceServer
}

func (s *server) IsAuthenticated(ctx context.Context, in *authorisation.AuthorisationRequest) (*authorisation.AuthorisationResponse, error) {

	log.Printf("Received: %v", in.GetJwt())
	return &authorisation.AuthorisationResponse{
		Authorised: false,
	}, nil

}

func main() {

	thor.LoadEnv()

	logger := logger.New("builder-api", os.Stdout, logger.LevelInfo, thor.GetTraceID, logger.DiscordEvent(logger.DiscordEventConfig{
		WebhookURL: "https://discord.com/api/webhooks/1224246134516744317/15uoW7BG5Jv6Ai1sMvqEU2a8an76vwflPzlDDKEOtTro43OEXlT24Trxr_PCisod9JiF",
		Error:      true,
	}))

	db := thor.GetPostgresClient()

	userService := user.NewUserServiceDB(db)
	userController := user.NewUserController(userService)

	authenticationService := authentication.NewAuthenticationService(userService)
	authenticationController := authentication.NewAuthenticationController(authenticationService)

	appModule := &common.Module{
		Controllers: []common.ControllerBase{
			userController,
			authenticationController,
		},
	}

	ctx := context.Background()

	c := core.NewNestFactory(thor.NewApp(ctx, thor.ThorConfig{
		Logger: logger,
	})).Create(appModule)

	go func() {
		lis, err := net.Listen("tcp", ":4444")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		s := grpc.NewServer()
		authorisation.RegisterAuthorisationServiceServer(s, &server{})

		log.Println("Running  GRPC SERVER")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	c.Listen(":8080")

}

package main

import (
	"auth/iam/authorisation"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:4444", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := authorisation.NewAuthorisationServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.IsAuthenticated(ctx, &authorisation.AuthorisationRequest{
		Jwt: "sdsdsd",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println(r.Authorised)
}

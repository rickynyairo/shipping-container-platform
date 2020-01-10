package main

import (
	"fmt"
	"github.com/micro/go-micro"
	pb "github.com/rickynyairo/shipping-container-platform/user-service/proto/user"
	"log"
)

func main() {

	// Creates a database connection and handles
	// closing it again before exit.
	db, err := CreateConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	// Automatically migrates the user struct
	// into database columns/types etc. This will
	// check for changes and migrate them each time
	// this service is restarted.
	db.AutoMigrate(&pb.User{})

	repo := &UserRepository{db}

	tokenService := &TokenService{repo}

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("user.service"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags.
	srv.Init()
	publisher := micro.NewPublisher("user.created", srv.Client())
	// pubsub := srv.Server().Options().Broker
	// Register handler
	pb.RegisterUserServiceHandler(srv.Server(), &Handler{repo, tokenService, publisher})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

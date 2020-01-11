package main

import (
	"errors"
	"fmt"

	"golang.org/x/net/context"

	// Import the generated protobuf code
	"log"
	"os"

	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	server "github.com/micro/go-micro/server"
	pb "github.com/rickynyairo/shipping-container-platform/consignment-service/proto/consignment"
	userProto "github.com/rickynyairo/shipping-container-platform/user-service/proto/user"
	vesselProto "github.com/rickynyairo/shipping-container-platform/vessel-service/proto/vessel"
)

const (
	port        = ":50051"
	defaultHost = "datastore:27017"
)

// var (
// 	srv micro.Service
// )

func main() {

	srv := micro.NewService(
		// the name of the service given in the protobuf definition
		micro.Name("go.micro.api.consignment"),
		micro.Version("latest"),
		micro.WrapHandler(AuthWrapper),
	)
	// parse command line flags
	srv.Init()
	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}
	client, err := CreateClient(uri)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.TODO())

	consignmentCollection := client.Database("shipping_container_platform").Collection("consignments")
	repository := &MongoRepository{consignmentCollection}
	vesselClient := vesselProto.NewVesselServiceClient("go.micro.api.vessel", srv.Client())

	h := &Handler{repository, vesselClient}

	// Register handler
	pb.RegisterShippingServiceHandler(srv.Server(), h)

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

// AuthWrapper is a high-order function which takes a HandlerFunc
// and returns a function, which takes a context, request and response interface.
// The token is extracted from the context set in our consignment-cli, that
// token is then sent over to the user service to be validated.
// If valid, the call is passed along to the handler. If not,
// an error is returned.
func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}
		if os.Getenv("DISABLE_AUTH") == "true" {
			return fn(ctx, req, resp)
		}

		// Note this is now uppercase (not entirely sure why this is...)
		token := meta["Token"]
		log.Println("Authenticating with token: ", token)

		// Auth here
		authClient := userProto.NewUserServiceClient("go.micro.api.user", client.DefaultClient)
		_, err := authClient.ValidateToken(context.Background(), &userProto.Token{
			Token: token,
		})
		if err != nil {
			return err
		}
		err = fn(ctx, req, resp)
		return err
	}
}

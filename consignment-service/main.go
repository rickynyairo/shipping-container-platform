package main

import (
	"context"
	"fmt"

	// Import the generated protobuf code
	pb "github.com/rickynyairo/shipping-container-platform/consignment-service/proto/consignment"
	vesselProto "github.com/rickynyairo/shipping-container-platform/vessel-service/proto/vessel"
	micro "github.com/micro/go-micro"
	"log"
	"os"
)


const ( 
	port = ":50051"
	defaultHost = "datastore:27017"

)

func main() {

	srv := micro.NewService(
		// the name of the service given in the protobuf definition
		micro.Name("consignment.service"),
	)
	// parse command line flags
	srv.Init()
	uri := os.Getenv("DB_HOST")
	if uri == ""{
		uri = defaultHost
	}
	client, err := CreateClient(uri)
	if err != nil{
		log.Panic(err)
	}
	defer client.Disconnect(context.TODO())

	consignmentCollection := client.Database("shipping_container_platform").Collection("consignments")
	repository := &MongoRepository{consignmentCollection}
	vesselClient := vesselProto.NewVesselServiceClient("vessel.service", srv.Client())

	h := &Handler{repository, vesselClient}

	// Register handler
	pb.RegisterShippingServiceHandler(srv.Server(), h)

	// Run the server
	if err := srv.Run(); err != nil{
		fmt.Println(err)
	}
}

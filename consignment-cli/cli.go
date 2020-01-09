package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	micro "github.com/micro/go-micro"
	pb "github.com/rickynyairo/shipping-container-platform/consignment-service/proto/consignment"
	pbVessel "github.com/rickynyairo/shipping-container-platform/vessel-service/proto/vessel"
)

const (
	address         = "localhost:50051"
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	service := micro.NewService(micro.Name("consignment.cli"))
	service.Init()

	client := pb.NewShippingServiceClient("consignment.service", service.Client())
	vesselClient := pbVessel.NewVesselServiceClient("vessel.service", service.Client())
	vessel := &pbVessel.Vessel{
		Capacity:  1000,
		MaxWeight: 5000,
		Name:      "Shawshank Redemption",
		Available: true,
	}
	_, err := vesselClient.Create(context.Background(), vessel)
	if err != nil {
		log.Fatalf("Could not create vessel: %v", err)
	}
	// Contact the server and print out its response.
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not create consignment: %v", err)
	}
	log.Printf("Created: %t", r)

	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}

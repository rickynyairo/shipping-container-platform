package main

import (
	"encoding/json"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"os"

	micro "github.com/micro/go-micro"
	metadata "github.com/micro/go-micro/metadata"
	pb "github.com/rickynyairo/shipping-container-platform/consignment-service/proto/consignment"
	pbVessel "github.com/rickynyairo/shipping-container-platform/vessel-service/proto/vessel"
)

const (
	address         = "localhost:50051"
	defaultFilename = "consignment.json"
	token           = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImlkIjoiMTNhM2Y1YTctNTYyYy00NzI3LWE1NmEtMDViNzYyY2M0ZTQ4IiwibmFtZSI6IlJpY2t5IE55YWlybyIsImNvbXBhbnkiOiJCQkMiLCJlbWFpbCI6InJpY2t5bnlhaXJvODlAZ21haWwuY29tIiwicGFzc3dvcmQiOiIkMmEkMTAkLjBQa09FRTVuZWk5ZUZKbzFSM0o3LllOMFdDZ3ZrNWpDeE9IcVpUL1BZLk9rNzAvV0ZkRDYifSwiZXhwIjoxNTc4ODc5MDYzLCJpc3MiOiJzaGlwcHkudXNlciJ9.SyaDpKOWvZExnDSj2YolUVEOyP891O8Xu4VWQqEvY5Y"
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
	service := micro.NewService(micro.Name("shippy.conscli"))
	service.Init()

	client := pb.NewShippingServiceClient("shippy.consignment", service.Client())
	vesselClient := pbVessel.NewVesselServiceClient("shippy.vessel", service.Client())
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
	// token := os.Args[1]
	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"token": token,
	})
	r, err := client.CreateConsignment(ctx, consignment)
	if err != nil {
		log.Fatalf("Could not create consignment: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetConsignments(ctx, &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}

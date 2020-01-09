// vessel service repository
package main

import (
	"context"
	pb "github.com/rickynyairo/shipping-container-platform/vessel-service/proto/vessel"
	mongo "go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type repository interface {
	FindAvailable(spec *pb.Specification) (*pb.Vessel, error)
	Create(vessel *pb.Vessel) error
}

type VesselRepository struct {
	collection *mongo.Collection
}

// FindAvailable - checks a specification against a map of vessels,
// if capacity and max weight are below a vessels capacity and max weight,
// then return that vessel.
func (repository *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	// filter:= bson.M{"capacity": bson.M{"$lte", spec.Capacity}}
	filter:= bson.M{"name": "Shawshank Redemption"}
	// filter := bson.D{{
	// 	"capacity",
	// 	bson.D{{
	// 		"$lte",
	// 		spec.Capacity,
	// 	}, {
	// 		"$lte",
	// 		spec.MaxWeight,
	// 	}},
	// }}
	var vessel *pb.Vessel
	if err := repository.collection.FindOne(context.TODO(), filter).Decode(&vessel); err != nil {
		return nil, err
	}
	return vessel, nil
}

// Create a new vessel
func (repository *VesselRepository) Create(vessel *pb.Vessel) error {
	_, err := repository.collection.InsertOne(context.TODO(), vessel)
	return err
}
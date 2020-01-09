package main

import (
	"context"
	"fmt"
	pb "github.com/rickynyairo/shipping-container-platform/consignment-service/proto/consignment"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository interface {
	Create(consignment *pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
}

// MongoRepository implementation
type MongoRepository struct {
	collection *mongo.Collection
}

// create consignment
func (repository *MongoRepository) Create(consignment *pb.Consignment) error {
	_, err := repository.collection.InsertOne(context.Background(), consignment)
	return err
}

// Get all
func (repository *MongoRepository) GetAll() ([]*pb.Consignment, error) {
	cur, err := repository.collection.Find(context.Background(), nil, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("Getting all consignments ==> ", cur)
	var consignments []*pb.Consignment
	for cur.Next(context.Background()) {
		var consignment *pb.Consignment
		if err := cur.Decode(&consignment); err != nil {
			return nil, err
		}
		consignments = append(consignments, consignment)
	}
	return consignments, err
}

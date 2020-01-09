// vessel service handler
package main

import (
	"context"
	pb "github.com/rickynyairo/shipping-container-platform/vessel-service/proto/vessel"
)


type Handler struct {
	repository
}

// FindAvailable vessels
func (s *Handler) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {

	// Find the next available vessel
	vessel, err := s.repository.FindAvailable(req)
	if err != nil {
		return err
	}

	// Set the vessel as part of the response message type
	res.Vessel = vessel
	return nil
}

// Create a new vessel
func (s *Handler) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	if err := s.repository.Create(req); err != nil {
		return err
	}
	res.Vessel = req
	return nil
}
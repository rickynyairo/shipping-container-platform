# Makefile for the consignment microservice

build:
	protoc -I. --go_out=plugins=grpc:. \
	  proto/consignment/consignment.proto

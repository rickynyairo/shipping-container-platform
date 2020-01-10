# Makefile for the consignment microservice

build:
	protoc -I. --go_out=plugins=micro:. \
	  proto/consignment/consignment.proto
	@echo 'Build Docker image'
	# docker build -t consignment-service .

run:
	docker run -p 50051:50051 consignment-service
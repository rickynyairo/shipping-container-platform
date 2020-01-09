# Makefile for the user microservice

build:
	protoc -I. --go_out=plugins=micro:. \
	  proto/user/user.proto
	@echo 'Build Docker image'
	docker build -t user-service .

run:
	docker run -p 50051:50051 user-service
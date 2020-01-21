# Makefile for the user microservice

build:
	protoc -I. --go_out=plugins=micro:. \
	  proto/user/user.proto
	@echo 'Build Docker image'
	docker build -t user-service .

run:
	docker run --net="host" \
		-p 50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		user-service
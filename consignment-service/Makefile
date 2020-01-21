
build:
	docker build -t consignment-service .

run:
	docker run --net="host" \
		-p 50053 \
		-e MICRO_SERVER_ADDRESS=:50053 \
		-e MICRO_REGISTRY=mdns \
		-e DB_HOST="mongodb://localhost:27017" \
		consignment-service
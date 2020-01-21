
build:
	docker build -t vessel-service .

run:
	docker run --net="host" \
		-p 50052 \
		-e MICRO_SERVER_ADDRESS=:50052 \
		-e MICRO_REGISTRY=mdns \
		-e DB_HOST="mongodb://localhost:27017" \
		vessel-service
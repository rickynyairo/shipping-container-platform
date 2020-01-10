# Makefile for the email microservice

build:
	@echo 'Build Docker image'
	docker build -t email-service .

run:
	docker run -p 50051:50051 email-service
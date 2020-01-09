build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t user-cli .

run:
	docker run user-cli
# vessel-service/Dockerfile
FROM golang:alpine as builder

RUN apk --no-cache add git

WORKDIR /app/vessel-service

COPY . .

# RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o vessel-service -a -installsuffix cgo main.go handler.go repository.go datastore.go


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/vessel-service .

CMD ["./vessel-service"]
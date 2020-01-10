FROM golang:alpine as builder

RUN apk --no-cache add git

WORKDIR /app/email-service

COPY . .

# RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o email-service


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/email-service .

CMD ["./email-service"]
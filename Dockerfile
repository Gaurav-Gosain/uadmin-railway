FROM golang:1.9-alpine as builder
LABEL maintainer="Gaurav Gosain"

RUN apk add --no-cache gcc musl-dev git

COPY . .

ENV GO111MODULE=off
ENV GOFLAGS=-mod=vendor

RUN go mod download

# Build the Go app
RUN go build .

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./railway"]
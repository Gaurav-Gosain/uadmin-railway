FROM golang:1.16.3-alpine3.13 AS builder
LABEL maintainer="Gaurav Gosain"

RUN apk add --no-cache gcc musl-dev git

COPY . ./workspace

WORKDIR /workspace

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

RUN go version

RUN go help

RUN go get -u -v github.com/uadmin/uadmin/...

# Build the Go app
RUN go build ./workspace

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./workspace/railway"]
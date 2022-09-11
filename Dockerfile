FROM golang:latest as builder
LABEL maintainer="Gaurav Gosain"

RUN apk add --no-cache gcc musl-dev git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o main .

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./main"]
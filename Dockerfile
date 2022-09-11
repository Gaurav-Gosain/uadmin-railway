FROM golang:1.9-alpine
LABEL maintainer="Gaurav Gosain"

RUN apk add --no-cache gcc musl-dev git

WORKDIR /app/server
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

# Build the Go app
RUN go build -v .

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["railway"]
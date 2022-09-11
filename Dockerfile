FROM golang

ENV GO111MODULE=on

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./railway"]
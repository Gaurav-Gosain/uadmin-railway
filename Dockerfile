FROM golang:1.18-alpine AS BUILDER
RUN apk add --no-cache gcc g++ git openssh-client
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOPROXY=https://goproxy.cn,direct \
    go build -ldflags="-w -s" -o server

RUN go mod download
RUN go build .
EXPOSE 8080

CMD [ "./railway" ]
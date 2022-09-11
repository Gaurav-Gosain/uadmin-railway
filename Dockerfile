FROM golang:1.9-alpine
LABEL maintainer="Gaurav Gosain"

RUN apk add --no-cache gcc musl-dev

RUN go mod download
RUN go build .
EXPOSE 8080

CMD [ "./railway" ]
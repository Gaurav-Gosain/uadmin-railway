FROM golang:1.9-alpine
LABEL maintainer="Gaurav Gosain"

RUN apk add --no-cache gcc musl-dev git

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["railway"]
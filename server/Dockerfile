FROM golang:1.21.4 as builder

# Set the working directory in the container.
WORKDIR /app

# Copy go mod and sum files.
COPY go.* ./

# Download all dependencies.
RUN go mod download

# Copy the source code into the container.
COPY . .

# Install the dependencies.
RUN go get github.com/99designs/gqlgen
RUN go get golang.org/x/tools/go/packages
RUN go get github.com/urfave/cli/v2

# Generate the GraphQL types.
RUN go run github.com/99designs/gqlgen generate

# Build the command inside the container.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o server ./cmd/main

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/
FROM alpine:3
RUN apk add --no-cache ca-certificates

# Copy the server executable from the builder image.
COPY --from=builder /app/server /server

# Document that the service listens on port 8080.
EXPOSE 8080

# Run the web service on container startup.
CMD ["/server"]
# Stage 1: Build the application
FROM golang:latest AS build

# Set the working directory inside the container to the root of the project
WORKDIR /go/src/github.com/yustinayasin/florist-gin

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project to the container
COPY . .

# Build the Go application
RUN go build -o main .

# Stage 2: Create a minimal runtime container
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /go/src/github.com/yustinayasin/florist-gin

# Copy the binary from the build stage
COPY --from=build /go/src/github.com/yustinayasin/florist-gin/main .

# Expose the port the application runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]

# Stage 1: Build the backend app
FROM golang:1.22-alpine AS build
# Set the working directory inside the container to the root of the project
WORKDIR /src
# Download dependencies
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download 
# Build the Go application
RUN -mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=.,target=. \
    go build -o /bin .

# Stage 2: Create a minimal runtime container
FROM alpine:latest
# Copy the binary from the build stage
COPY --from=build /bin .
# Expose the port the application runs on
EXPOSE 8080
# Command to run the executable
CMD ["./main"]

# Start from golang base image
FROM golang:alpine as builder

# Set the current working directory inside the container
WORKDIR /app
COPY ./fixture/ports.json .
ARG file= /usr/src/app/fixture/ports.json

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build the Go app
RUN cd cmd && go build -o app

# Start a new stage from scratch
FROM alpine:latest

# Add certificates so we can validate TLS certificates of external services.
RUN apk --no-cache update && \
      apk --no-cache add ca-certificates && \
      rm -rf /var/cache/apk/*

# Copy the Pre-built binary file from the previous stage.
COPY --from=builder app .

#Command to run the executable
CMD ./cmd/app

# Start from golang base image
FROM golang:1.19-alpine as builder

# Add Maintainer info
LABEL maintainer="Benson Macharia"

# Set the current working directory inside the container 
WORKDIR /app

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Build the Go app
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo

# Start a new stage from scratch
FROM alpine:latest

WORKDIR /golang/

# Copy the Pre-built binary file from the previous stage. 
COPY --from=builder /app/book_store_api .
COPY --from=builder /app/.env .       

# Expose port 8080 to the outside world. Must match with the one in DBconfig.
EXPOSE 8080

#Command to run the executable
CMD ["./book_store_api"]
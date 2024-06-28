# Use an official Golang image as a parent image
FROM golang:1.22.3-alpine

# Set the Current Working Directory inside the container
# /app klasorunu docker icin olusturuyor
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Install build dependencies (if needed)
RUN apk add --no-cache build-base
##RUN apk add --no-cache gcc musl-dev

# Copy the source code into the container
COPY . .

# Build the Go app
# chat.go yu da dahil eder main adinda tek bir executable olusturuyor
ENV CGO_ENABLED=1 
RUN go build -o main .
#RUN CGO_ENABLED=1 go build -o main .

# Expose port 8065 to the outside world
EXPOSE 8065

# Command to run the executable
CMD ["./main"]
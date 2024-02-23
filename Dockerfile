FROM golang:latest

# Set the working directory in the container
WORKDIR /go/src/app

# Copy the local package files to the container's workspace
COPY . .

RUN go build -o main .

# Expose the port the application runs on
EXPOSE 8080

# Run the Go application
CMD ["./main"]
FROM golang:1.20.5 

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

COPY main.go .

# Download all dependencies. 
RUN go mod download

COPY . .

# Build the Go app
RUN go build -o bin .

# Expose port 8080 to the outside world
EXPOSE 8000

# Command to run 
CMD ["./bin"]
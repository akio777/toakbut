# Use the official Golang image as the base image
FROM golang:1.21.4

RUN apt-get update && apt-get install -y tzdata
ENV TZ=Asia/Bangkok
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# Set the working directory inside the container
WORKDIR /build

# Copy the Go Modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code to the container
COPY . .

# Build the Go application
RUN go build -o toakbut ./cmd/toakbut/main.go

# Command to run the executable
CMD ["./toakbut"]

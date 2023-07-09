# Use the official Go image as the base image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the necessary files into the container
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the project files
COPY . .

# Build and install the Go application
RUN go build -o mpc_wallet

# Expose the desired port for your Go application
EXPOSE 8080

# Specify the command to run your Go application
CMD ["./mpc_wallet"]

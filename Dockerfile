# Start from a Golang base image
FROM golang:latest

# Set the working directory
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Install any necessary dependencies
RUN go mod download

# Expose the port the app will be running on
EXPOSE 8080

# Set environment variables
ENV ENVIRONMENT=development

# Build the app
RUN go build -o main .

# Start the app
CMD ["/app/main"]

# Start from a base Go image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY . .

# Compile the Go application
RUN go build -o main .

# Specify the command to run the compiled binary
CMD ["./main"]

# Use an official Go runtime as a parent image
FROM golang:1.22.1 as builder

# Set an argument for the service name
ARG SERVICE
ARG PORT
# Create a directory for the shared packages
RUN mkdir /pkg
COPY pkg/ /pkg/

# Create a directory for internal packages
RUN mkdir /internal
COPY internal/ /internal/

# Set the working directory to the service directory
WORKDIR /internal/${SERVICE}
RUN go list -m all
# Download module dependencies
RUN go mod download

# Build the application and output the executable to /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /${SERVICE} cmd/main.go
# Set the command to run the application
FROM alpine:3.15
ARG SERVICE
ARG PORT
WORKDIR /app
# copy from builder the binary
COPY --from=builder /${SERVICE} ./myapp

EXPOSE ${PORT}
CMD ["./myapp"]

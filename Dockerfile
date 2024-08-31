FROM golang:1.23 as build

WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download the Go module dependencies
RUN go mod download

COPY . .
RUN go build -o /myapp ./cmd/main.go
FROM alpine:latest as run

# Copy the application executable from the build image
COPY --from=build /myapp /myapp
WORKDIR /app
EXPOSE 8080
CMD ["/myapp"]

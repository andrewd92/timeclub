FROM golang:1.17.1-alpine AS builder
# All these steps will be cached
RUN mkdir /app
WORKDIR /app
## Copy go.mod and go.sum files to workspace
COPY go.mod .
COPY go.sum .

# Copy the source code
COPY . .

# Get dependencies, will also be cached if you don't change mod/sum
RUN go mod download

# Build the binary
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /go/bin/club_service

#FINAL IMAGE
FROM scratch

ENV VIPER_CONFIG_NAME=config

# Copy our static executable.
COPY --from=builder /app/config/docker.yml /docker.yml
COPY --from=builder /go/bin/club_service /go/bin/club_service
# CMD
ENTRYPOINT  ["/go/bin/club_service"]
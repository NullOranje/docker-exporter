## Compile phase
FROM golang:alpine AS builder

## Install Git to all 'go get' to work
RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/docker-exporter/
COPY . .

# Go get stuff
RUN go get -d -v

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /docker-exporter

## Container build
FROM scratch

# Copy the binary
COPY --from=builder /docker-exporter /docker-exporter

# Run the binary
ENTRYPOINT ["/docker-exporter"]

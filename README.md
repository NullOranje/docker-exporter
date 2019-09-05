# docker-exporter

Simple application to expose Docker container running state as Prometheus metrics.

By default, exposes metrics on all interfaces, port `12345`

## Building
### As standalone
From the source directory, run `go build`

### As container 
From the source directory, run `docker build -t docker-exporter .`

## Running
### As standalone
No configuration is required; run the binary on the local machine.

### As container
There are two key things here: mounting the Docker daemon socket to the container and exposing the Prometheus metrics endpoint 
package main

import (
	"docker-exporter/exporter"
	"log"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	options := types.ContainerListOptions{All: true}
	containerCh := make(chan exporter.Update, 10)
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	go exporter.DockerContainerMetrics(containerCh)
	go func() {
		for {
			containers, err := cli.ContainerList(ctx, options)
			if err != nil {
				log.Println(err)
			}

			for _, container := range containers {
				for _, name := range container.Names {
					containerCh <- exporter.Update{
						State:    container.State,
						Name:     name,
						Instance: hostname,
					}
				}
			}

			time.Sleep(1 * time.Second)
		}
	}()

	err = exporter.HandleHTTP(":12345")
	if err != nil {
		log.Println(err)
	}
}

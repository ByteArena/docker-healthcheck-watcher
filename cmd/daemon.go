package main

import (
	"context"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func onContainerHealthCheckFailure(service string) {
	log.Println("FAILURE CONTAINER", service)
}

func main() {
	cli, clientError := client.NewEnvClient()
	ctx := context.Background()

	if clientError != nil {
		log.Panicln(clientError)
	}

	stream, err := cli.Events(ctx, types.EventsOptions{})

	for {
		select {
		case msg := <-err:
			log.Panicln(msg)
		case msg := <-stream:
			if msg.Action == "health_status: unhealthy" {
				onContainerHealthCheckFailure(msg.Actor.Attributes["image"])
			}
		}
	}
}

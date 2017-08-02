package main

import (
	"context"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	"github.com/bytearena/docker-healthcheck-watcher/integration"
	"github.com/bytearena/docker-healthcheck-watcher/template"
	t "github.com/bytearena/docker-healthcheck-watcher/types"
)

func onContainerDieFailure(service string, exitCode string) {
	errorMessage := t.ErrorMessage{
		ServiceName:   service,
		ServiceStatus: "died (exited with code " + exitCode + ")",
		Log:           "",
	}

	message := template.MakeTemplate(errorMessage)

	output := slack.Publish(message)

	log.Println(service, "failure, sent message", output)
}

func onContainerHealthCheckFailure(service string) {
	errorMessage := t.ErrorMessage{
		ServiceName:   service,
		ServiceStatus: "unhealthy (running)",
		Log:           "",
	}

	message := template.MakeTemplate(errorMessage)

	output := slack.Publish(message)

	log.Println(service, "failure, sent message", output)
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

			if msg.Action == "die" {
				exitCode := msg.Actor.Attributes["exitCode"]

				if exitCode != "0" {
					onContainerDieFailure(msg.Actor.Attributes["image"], exitCode)
				}
			}
		}
	}
}

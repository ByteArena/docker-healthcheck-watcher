# docker-healthcheck-watcher

> Monitor container healthchecks

<img src="./.github/demo.png" />

### The issue

If you are using a micro-service oriented architecture and you want to make sure all your services are up and running, you need some sort of application healthchecks.

Usually each service must expose an HTTP interface (something like /health). In a web application it's easy, you can just create a new route responding to the healthcheck but how would you do for system binaries for example?

### The solution

You can specify a `HEALTHCHECK` instruction in your Dockerfile. Docker will be executing the command at a regular time (interval and timeout are configurable).
Almost no monitoring/alerting tool uses that Docker feature, so we decided to build our own.

The checkhealth doesn't pass the boundaries of the container, there is no security or disclosure issue. Even for non HTTP service likes batches you can provide a custom command (bash script) to check the health of your program.

If Docker detects an unhealthy service, it will send a message to our tool and we can decide what to do (see integrations).

### Integrations

At the moment only Slack is supported for alerting. The configuration is the following:

```yml
# Slack endpoint with the authentication token
SLACK_URL: "https://hooks.slack.com/services/[...]"

# Channel, private group, or IM channel to send message to.
SLACK_CHANNEL: "#ops"

# Set your bot's user name.
SLACK_USERNAME: "dockerwatcher"

# Emoji to use as the icon for this message
SLACK_ICONEMOJI: ":robot_face:"
```

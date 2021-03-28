package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

const (
	usage = `Usage: %s [options...] <command> [service]

Options:
`
	commands = `
Commands:
  build              Build or rebuild services
  config             Validate and view the Compose file
  create             Create services
  down               Stop and remove containers, networks, images, and volumes
  events             Receive real time events from containers
  exec               Execute a command in a running container
  help               Get help on a command
  images             List images
  kill               Kill containers
  log, logs          View output (stdout+stderr) from services
  pause              Pause services
  port               Print the public port for a port binding
  ps                 List services
  pull               Pull remote repositories or service images
  push               Push remote repositories or service images
  restart            Restart services
  rm                 Remove stopped services
  run                Run a one-off command
  scale              Set number of containers for a service
  start              Start services
  stop               Stop services
  top                Display the running processes
  unpause            Unpause services
  up                 Create and start containers
  version            Show version information and quit
`
)

func usageAndExit(msg string) {
	if msg != "" {
		fmt.Fprintf(os.Stderr, msg)
		fmt.Fprintf(os.Stderr, "\n\n")
	}
	pflag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}

func errExit(err error) {
	fmt.Fprintf(os.Stderr, err.Error())
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}

package main

import (
	"fmt"
	"os"

	"github.com/multi-repo-toolbox/dockerless-compose/command"
	"github.com/multi-repo-toolbox/dockerless-compose/command/up"
	"github.com/multi-repo-toolbox/dockerless-compose/command/version"
	"github.com/multi-repo-toolbox/dockerless-compose/config"
	"github.com/spf13/pflag"
)

var (
	RepoName string
	GitHash  string
	Version  string
	BuildAt  string
)

var (
	help       = pflag.BoolP("help", "h", false, "Show general help or about a command")
	configFile = pflag.StringP("file", "f", "./dockerless.yaml", "Specify an alternative compose file")
)

func main() {
	pflag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage, os.Args[0])
		pflag.PrintDefaults()
		fmt.Fprintf(os.Stderr, commands)
	}
	pflag.CommandLine = pflag.NewFlagSet(os.Args[0], pflag.ContinueOnError)
	pflag.Parse()
	if pflag.NArg() < 1 {
		usageAndExit("select a command")
	}
	config.Setup(*configFile)
	allCommands()
	if *help {
		command.Help(pflag.Arg(0))
		return
	}
	command.Process(pflag.Arg(0), os.Args)
}

func allCommands() {
	command.Register(new(version.Command))
	command.Register(new(up.Command))
}

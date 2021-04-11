package command

import (
	"github.com/multi-repo-toolbox/dockerless-compose/config"
	"github.com/spf13/pflag"
)

type Commander interface {
	Name() string                           // declares a name of command in CLI
	Desc() string                           // briefly describes usage of a command
	Flags() *pflag.FlagSet                  // returns flags used for a command
	Config() (*config.Configuration, error) // loads file configuration for a command
	Action(*config.Configuration) error     // executes action of the command
}

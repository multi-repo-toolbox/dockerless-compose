package command

import "github.com/spf13/pflag"

type Commander interface {
	Name() string          // declares a name of command in CLI
	Desc() string          // briefly describes usage of a command
	Flags() *pflag.FlagSet // returns flags used for a command
	Action() error         // executes action of the command
}

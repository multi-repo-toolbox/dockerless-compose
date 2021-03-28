package command

import "flag"

type Commander interface {
	Name() string
	Desc() string
	Flags() *flag.FlagSet
	Action() error
}

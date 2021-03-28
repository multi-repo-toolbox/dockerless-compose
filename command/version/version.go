package version

import (
	"flag"
	"fmt"
)

var (
	AppName string
	Version string
	GitHash string
	BuildAt string
)

type Command struct{}

func (*Command) Name() string {
	return "version"
}

func (*Command) Desc() string {
	return "Show version information and quit"
}

func (*Command) Flags() *flag.FlagSet {
	return nil
}

func (c *Command) Action() error {
	fmt.Println(AppName, "version", Version)
	return nil
}

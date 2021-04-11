package version

import (
	"fmt"

	"github.com/multi-repo-toolbox/dockerless-compose/config"
	"github.com/spf13/pflag"
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

func (*Command) Flags() *pflag.FlagSet {
	return nil
}

func (*Command) Config() (*config.Configuration, error) {
	return nil, nil
}

func (c *Command) Action(*config.Configuration) error {
	fmt.Println(AppName, "version", Version)
	return nil
}

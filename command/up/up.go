package up

import "github.com/spf13/pflag"

type Command struct {
	detached *bool
}

func (*Command) Name() string {
	return "up"
}

func (*Command) Desc() string {
	return "Create and start applications/containers"
}

func (c *Command) Flags() *pflag.FlagSet {
	fs := pflag.NewFlagSet(c.Name(), pflag.ExitOnError)
	c.detached = fs.BoolP("detached", "d", false, "Detach applications/containers")
	return fs
}

func (c *Command) Action() error {
	if *c.detached {
		print("XXX detached")
	}
	print("XXX implement it")
	return nil
}

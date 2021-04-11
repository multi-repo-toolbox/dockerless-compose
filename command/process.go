package command

import (
	"errors"
	"fmt"
)

var registry = make(map[string]Commander)
var errCmdNotFound = errors.New("command not found")

func Process(name string, args []string) error {
	if c, ok := registry[name]; ok {
		fs := c.Flags()
		if fs != nil {
			fs.Parse(args)
		}
		cfg, err := c.Config()
		if err != nil {
			return err
		}
		return c.Action(cfg)
	}
	return errCmdNotFound
}

func Help(name string) {
	if c, ok := registry[name]; ok {
		fmt.Println(c.Desc())
	}
}

func Register(c Commander) {
	registry[c.Name()] = c
}

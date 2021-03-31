package command

import "fmt"

var registry = make(map[string]Commander)

func Process(name string, args []string) {
	//	name := args[0]
	if c, ok := registry[name]; ok {
		fs := c.Flags()
		if fs != nil {
			fs.Parse(args)
		}
		c.Action()
	}
}

func Help(name string) {
	if c, ok := registry[name]; ok {
		fmt.Println(c.Desc())
	}
}

func Register(c Commander) {
	registry[c.Name()] = c
}

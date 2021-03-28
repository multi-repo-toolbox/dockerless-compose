package command

import "fmt"

var registry = make(map[string]Commander)

func Process(name string) {
	if c, ok := registry[name]; ok {
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

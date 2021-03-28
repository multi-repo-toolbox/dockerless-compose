package config

import (
	"os"

	"github.com/multi-repo-toolbox/config"
)

var (
	cfg   config.Extended
	fname string
	file  *os.File
)

func Setup(f string) {
	fname = f
}

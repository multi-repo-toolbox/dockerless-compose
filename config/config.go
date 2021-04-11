package config

import (
	"os"

	"github.com/multi-repo-toolbox/fileconf"
)

type Configuration struct {
	fileconf.Extended
}

var (
	cfg   *fileconf.Extended
	fname string
	file  *os.File
)

// Setup set file name of the configuration file. It could not check
// the file availability.
func SetName(file string) {
	fname = file
}

// LoadFile loads and parses applications configuration from the file.
func LoadFile() (*Configuration, error) {
	var err error
	if cfg, err = fileconf.LoadFile(fname); err != nil {
		return nil, err
	}
	return &Configuration{*cfg}, nil
}

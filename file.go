package lamend

import (
	"errors"
	"os"
	"path/filepath"
)

const (
	extYml  = ".yml"
	extYaml = ".yaml"
)

var errUnsupportedFileFormat = errors.New("unsupported file format")

func Read(path string) ([]byte, error) {
	if ext := filepath.Ext(path); ext != extYml && ext != extYaml {
		return nil, errUnsupportedFileFormat
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return b, nil
}

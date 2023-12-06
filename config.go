package lamend

import (
	"errors"

	"github.com/goccy/go-yaml"
)

type Config struct {
	FunctionName string `yaml:"function_name"`
	ImageURI     string `yaml:"image_uri"`
}

var errEmptyRequiredField = errors.New("function_name or image_uri is empty")

func UnmarshalConfigYAML(b []byte) (*Config, error) {
	var cfg Config
	err := yaml.Unmarshal(b, &cfg)
	if err != nil {
		return nil, err
	}
	if ok := isRequiredFieldEmpty(cfg); ok {
		return nil, errEmptyRequiredField
	}
	return &cfg, nil
}

func isRequiredFieldEmpty(cfg Config) bool {
	return cfg.FunctionName == "" || cfg.ImageURI == ""
}

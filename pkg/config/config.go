package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Lowercase      bool     `yaml:"lowercase"`
	Latin          bool     `yaml:"latin"`
	SpecialChars   bool     `yaml:"special_chars"`
	SensitiveData  bool     `yaml:"sensitive_data"`
	SensitiveWords []string `yaml:"sensitive_words"`
}

func Load(path string) (*Config, error) {
	cfg := Config{
		Lowercase:      true,
		Latin:          true,
		SpecialChars:   true,
		SensitiveData:  true,
		SensitiveWords: make([]string, 0),
	}

	if path == "" {
		return &cfg, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return &cfg, nil
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

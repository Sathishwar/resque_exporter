package resqueExporter

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	GuardIntervalMillis int64        `yaml:"guard_interval_millis"`
	ResqueNamespace     *ResqueNamespaceConfig       `yaml:"resque_namespace"`
	Redis               *RedisConfig `yaml:"redis"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int64  `yaml:"db"`
}

type ResqueNamespaceConfig struct {
	Namespace []string `yaml:"namespace"`
}

func loadConfig(configPath string) (*Config, error) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("Failed to load config; path:<%s>, err:<%s>", configPath, err)
	}

	var config *Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("Failed to parse yaml; err:<%s>", err)
	}

	return config, nil
}

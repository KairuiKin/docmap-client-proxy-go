// internal/config/config.go
package config

import (
	"docmap-client-proxy-go/internal/common"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

// Config 用于映射我们的 YAML 文件结构
type Config struct {
	Tunnel struct {
		ServerURL string `yaml:"server_url"`
		Address   string `yaml:"address"`
	} `yaml:"tunnel"`

	Classification struct {
		AllowedSecret []struct {
			Level string `yaml:"level"`
			Code  string `yaml:"code"`
			Name  string `yaml:"name"`
		} `yaml:"allowed_secret"`
		DefaultSecret struct {
			Level string `yaml:"level"`
			Code  string `yaml:"code"`
			Name  string `yaml:"name"`
		} `yaml:"default_secret"`
	} `yaml:"classification"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	common.Logger.Info("Configuration loaded from %s", path)
	return &cfg, nil
}

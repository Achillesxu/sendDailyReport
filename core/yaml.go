package core

import (
	"bytes"
	"github.com/go-yaml/yaml"
)

type SmtpConf struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// ConfigData is our common data struct
type ConfigData struct {
	Smtp SmtpConf `yaml:"smtp_config"`
}

// ToYAML dumps the ConfigData struct to
// a YAML format bytes.Buffer
func (t *ConfigData) ToYAML() (*bytes.Buffer, error) {
	d, err := yaml.Marshal(t)
	if err != nil {
		return nil, err
	}

	b := bytes.NewBuffer(d)

	return b, nil
}

// Decode will decode into TOMLData
func (t *ConfigData) Decode(data []byte) error {
	return yaml.Unmarshal(data, t)
}

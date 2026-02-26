package config

import "github.com/pkg/errors"

type Config interface {
	Setting() Setting
}

type DefaultConfig struct {
	setting Setting
}

func (c *DefaultConfig) Setting() Setting {
	return c.setting
}

// New builds a config instance.
func New(setting Setting) (Config, error) {
	if err := setting.validate(); err != nil {
		return nil, errors.Wrap(err, "invalid setting")
	}
	return &DefaultConfig{setting: setting}, nil
}

// Mock returns config for tests.
func Mock() Config {
	return &DefaultConfig{setting: NewSetting()}
}

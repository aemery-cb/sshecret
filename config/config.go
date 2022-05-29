package config

import (
	"bytes"
	"errors"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	ConfigPath     string
	AuthorizedKeys map[string]string
}

func (c *Config) Write() error {
	var firstBuffer bytes.Buffer

	err := toml.NewEncoder(&firstBuffer).Encode(c.AuthorizedKeys)
	if err != nil {
		return err
	}
	err = os.WriteFile(c.ConfigPath, firstBuffer.Bytes(), 0777) //TODO: not this
	if err != nil {
		return err
	}
	return nil

}

func (c *Config) Read() error {
	if _, err := os.Stat(c.ConfigPath); errors.Is(err, os.ErrNotExist) {
		log.Panic("No config file found")
	}

	bytes, err := os.ReadFile(c.ConfigPath)
	if err != nil {
		return err
	}
	var keys map[string]string
	_, err = toml.Decode(string(bytes), &keys)
	if err != nil {
		return err
	}
	c.AuthorizedKeys = keys
	return nil
}
func NewConfigFrom(configPath string) (*Config, error) {
	conf := Config{
		ConfigPath: configPath,
	}

	if err := conf.Read(); err != nil {
		return nil, err
	}

	return &conf, nil

}

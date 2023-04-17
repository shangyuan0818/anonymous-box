package config

import "github.com/kelseyhightower/envconfig"

func NewEnvConfig[T any](prefix string, out T) func() (*T, error) {
	if err := envconfig.Process(prefix, &out); err != nil {
		return func() (*T, error) {
			return nil, err
		}
	}

	return func() (*T, error) {
		return &out, nil
	}
}

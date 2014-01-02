package gonfig

import (
	"os"
	"strings"
)

type EnvConfig struct {
	Configurable
	Prefix string
}

func NewEnvConfig(prefix string) ReadableConfig {
	cfg := &EnvConfig{NewMemoryConfig(), prefix}
	cfg.Load()
	return cfg
}

func (self *EnvConfig) Get(key string) interface{} {
	return self.Configurable.Get(self.Prefix + key)
}

func (self *EnvConfig) Load() (err error) {
	env := os.Environ()
	for _, pair := range env {
		kv := strings.Split(pair, "=")
		if kv != nil && len(kv) >= 2 {
			self.Set(strings.Replace(kv[0], self.Prefix, "", 1), kv[1])
		}
	}
	return nil
}

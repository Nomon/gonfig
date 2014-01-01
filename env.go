package gonfig

import (
	"os"
	"strings"
)

type EnvConfig struct {
	*MemoryConfig
	prefix string
}

func NewEnvConfig(prefix string) Configurable {
	cfg := &EnvConfig{&MemoryConfig{}, prefix}
	cfg.Load()
	return cfg
}

func (self *EnvConfig) Get(key string) interface{} {
	if self.data == nil {
		self.initialize()
	}
	return self.data[self.prefix+key]
}

func (self *EnvConfig) Load() (err error) {
	self.initialize()
	env := os.Environ()
	for _, pair := range env {
		kv := strings.Split(pair, "=")
		if kv != nil && len(kv) >= 2 {
			self.Set(strings.Replace(kv[0], self.prefix, "", 1), kv[1])
		}
	}
	return nil
}

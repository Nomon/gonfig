package gonfig

import (
	"os"
	"strings"
)

type envConfig struct {
	*MemoryConfig
	prefix string
}

func NewEnvConfig(prefix string) Configurable {
	cfg := &envConfig{&MemoryConfig{}, prefix}
	cfg.Load()
	return cfg
}

func (self *envConfig) Get(key string) interface{} {
	if self.data == nil {
		self.initialize()
	}
	return self.data[self.prefix+key]
}

func (self *envConfig) Load() (err error) {
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

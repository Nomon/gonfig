package gonfig

import (
	"os"
	"strings"
)

// EnvConfig can be used to read values from the environment
// into the underlaying Configurable
type EnvConfig struct {
	Configurable
	Prefix string
}

// Creates a new Env config backed by a memory config
func NewEnvConfig(prefix string) ReadableConfig {
	cfg := &EnvConfig{NewMemoryConfig(), prefix}
	cfg.Load()
	return cfg
}

// Loads the data from os.Environ() to the underlaying Configurable.
// if a Prefix is set then variables are imported with self.Prefix removed from the name
// so MYAPP_test=1 exported in env and read from ENV by EnvConfig{Prefix:"MYAPP_"} can be found from
// EnvConfig.Get("test")
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

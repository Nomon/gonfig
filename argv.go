package gonfig

import (
	"flag"
	"strings"
)

type ArgvConfig struct {
	Configurable
	Prefix string
}

func NewArgvConfig(prefix string) Configurable {
	cfg := &ArgvConfig{NewMemoryConfig(), prefix}
	cfg.Load()
	return cfg
}

func (self *ArgvConfig) Load() (err error) {
	flag.VisitAll(func(f *flag.Flag) {
		name := f.Name
		if self.Prefix != "" && strings.HasPrefix(f.Name, self.Prefix) {
			name = strings.Replace(name, self.Prefix, "", 1)
		}
		if getter, ok := f.Value.(flag.Getter); ok {
			self.Set(name, getter.Get())
		}
	})
	return nil
}

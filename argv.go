package gonfig

import (
	"flag"
	"strings"
)

type ArgvConfig struct {
	*MemoryConfig
	prefix string
}

func NewArgvConfig(prefix string) Configurable {
	cfg := &ArgvConfig{&MemoryConfig{}, prefix}
	cfg.Load()
	return cfg
}

func (self *ArgvConfig) Get(key string) interface{} {
	var f *flag.Flag
	if self.prefix != "" {
		f = flag.Lookup(self.prefix + key)
	} else {
		f = flag.Lookup(key)
	}
	if f == nil {
		return nil
	}
	if getter, ok := f.Value.(flag.Getter); ok {
		return getter.Get()
	}
	return nil
}

func (self *ArgvConfig) All() map[string]interface{} {
	data := make(map[string]interface{})
	flag.VisitAll(func(f *flag.Flag) {
		name := f.Name
		if self.prefix != "" && strings.HasPrefix(f.Name, self.prefix) {
			name = strings.Replace(name, self.prefix, "", 1)
		}
		if getter, ok := f.Value.(flag.Getter); ok {
			data[name] = getter.Get()
		}
	})
	return data
}

func (self *ArgvConfig) Load() (err error) {
	self.initialize()
	return nil
}

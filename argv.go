package gonfig

import (
	"flag"
	"strings"
)

type argvConfig struct {
	*memoryConfig
	prefix string
}

func NewArgvConfig(prefix string) Configurable {
	cfg := &argvConfig{&memoryConfig{}, prefix}
	cfg.Load()
	return cfg
}

func (self *argvConfig) Get(key string) interface{} {
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

func (self *argvConfig) All() map[string]interface{} {
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

func (self *argvConfig) Load() (err error) {
	self.initialize()
	return nil
}

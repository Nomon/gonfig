package gonfig

import (
	"encoding/json"
	"io/ioutil"
)

type jsonConfig struct {
	*MemoryConfig
	path string
}

func NewJsonConfig(path string) Configurable {
	cfg := &jsonConfig{&MemoryConfig{}, path}
	cfg.Load()
	return cfg
}

func (self *jsonConfig) Load() (err error) {
	self.initialize()
	var data []byte = make([]byte, 1024)
	if data, err = ioutil.ReadFile(self.path); err != nil {
		return err
	}
	out := make(map[string]interface{})
	if err := json.Unmarshal(data, &out); err != nil {
		return err
	}
	self.data = out
	return nil
}

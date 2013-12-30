package gonfig

import (
	"encoding/json"
	"io/ioutil"
)

type jsonConfig struct {
	*memoryConfig
	path string
}

func NewJsonConfig(path string) Configurable {
	cfg := &jsonConfig{&memoryConfig{}, path}
	cfg.Load()
	return cfg
}

func (self *jsonConfig) Load() (err error) {
	if self.data == nil {
		self.initialize()
	}
	var data []byte = make([]byte, 1024)
	if data, err = ioutil.ReadFile(self.path); err != nil {
		return err
	}
	out, err := self.unmarshal(data)
	if err != nil {
		return err
	}
	self.data = out
	return nil
}

func (self *jsonConfig) Save() (err error) {
	if self.data == nil {
		self.initialize()
	}
	b, err := json.Marshal(self.data)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(self.path, b, 0600); err != nil {
		return err
	}

	return nil
}

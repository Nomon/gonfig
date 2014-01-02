package gonfig

import (
	"encoding/json"
	"io/ioutil"
)

type JsonConfig struct {
	Configurable
	Path string
}

func unmarshalJson(bytes []byte) (map[string]interface{}, error) {
	out := make(map[string]interface{})
	if err := json.Unmarshal(bytes, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Returns a new Configurable backed by a json file at path.
// The file does not need to exist, if it does not exist the first Save call will create it.
func NewJsonConfig(path string) Configurable {
	cfg := &JsonConfig{NewMemoryConfig(), path}
	cfg.Load()
	return cfg
}

func (self *JsonConfig) Load() (err error) {
	var data []byte = make([]byte, 1024)
	if data, err = ioutil.ReadFile(self.Path); err != nil {
		return err
	}
	out, err := unmarshalJson(data)
	if err != nil {
		return err
	}
	self.Configurable.Reset(out)
	return nil
}

func (self *JsonConfig) Save() (err error) {
	b, err := json.Marshal(self.Configurable.All())
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(self.Path, b, 0600); err != nil {
		return err
	}

	return nil
}

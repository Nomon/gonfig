package gonfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type JsonConfig struct {
	Configurable
	Path string
}

func unmarshalJson(bytes []byte) (map[string]string, error) {
	out := make(map[string]interface{})
	if err := json.Unmarshal(bytes, &out); err != nil {
		return nil, err
	}
	var output map[string]string = make(map[string]string)

	for k, v := range out {
		switch v := v.(type) {
		case int:
			output[k] = fmt.Sprintf("%d", v)
		case float64:
			output[k] = fmt.Sprintf("%f", v)
		case string:
			output[k] = v
		case bool:
			if v {
				output[k] = "true"
			} else {
				output[k] = "false"
			}
		default:
			// i isn't one of the types above
		}
	}
	return output, nil
}

// Returns a new WritableConfig backed by a json file at path.
// The file does not need to exist, if it does not exist the first Save call will create it.
func NewJsonConfig(path string, cfg ...Configurable) WritableConfig {
	if len(cfg) == 0 {
		cfg = append(cfg, NewMemoryConfig())
	}
	LoadConfig(cfg[0])
	conf := &JsonConfig{cfg[0], path}
	LoadConfig(conf)
	return conf
}

// Attempts to load the json configuration at JsonConfig.Path
// and Set them into the underlaying Configurable
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

// Attempts to save the configuration from the underlaying Configurable to json file at JsonConfig.Path
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

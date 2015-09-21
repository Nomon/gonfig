package gonfig

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"io/ioutil"
)

type JsonConfig struct {
	Configurable
	Path string
}

func unmarshalJsonSegment(jsonSegment map[string]interface{}, segmentPath string, output map[string]string) {
	if segmentPath != "" {
		segmentPath += ":"
	}

	for k, v := range jsonSegment {
		keyWithPath := segmentPath + k

		switch v := v.(type) {
		case map[string]interface{}:
			unmarshalJsonSegment(v, keyWithPath, output)
		case []interface{}:
			var buffer bytes.Buffer
			for _, sVal := range v {
				buffer.WriteString(fmt.Sprintf("%v,", sVal))
			}
			output[keyWithPath] = strings.Trim(buffer.String(), ",")
		default:
			output[keyWithPath] = fmt.Sprintf("%v", v)
		}
	}
}

func unmarshalJson(bytes []byte) (map[string]string, error) {
	out := make(map[string]interface{})
	if err := json.Unmarshal(bytes, &out); err != nil {
		return nil, err
	}

	var output map[string]string = make(map[string]string)
	unmarshalJsonSegment(out, "", output)

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

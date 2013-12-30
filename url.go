package gonfig

import (
	"io/ioutil"
	"net/http"
)

type urlConfig struct {
	*memoryConfig
	url string
}

// Returns a new Configurable backed by JSON at url
func NewUrlConfig(url string) *urlConfig {
	return &urlConfig{&memoryConfig{}, url}
}

func (self *urlConfig) Load() error {
	if self.data == nil {
		self.initialize()
	}
	resp, err := http.Get(self.url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	out, err := self.unmarshal(body)
	if err != nil {
		return err
	}
	self.data = out
	return nil
}

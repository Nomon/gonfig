package gonfig

import (
	"io/ioutil"
	"net/http"
)

type UrlConfig struct {
	*MemoryConfig
	url string
}

// Returns a new Configurable backed by JSON at url
func NewUrlConfig(url string) *UrlConfig {
	return &UrlConfig{&MemoryConfig{}, url}
}

func (self *UrlConfig) Load() error {
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

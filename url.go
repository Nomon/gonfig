package gonfig

import (
	"io/ioutil"
	"net/http"
)

type UrlConfig struct {
	Configurable
	url string
}

// Returns a new Configurable backed by JSON at url
func NewUrlConfig(url string) ReadableConfig {
	return &UrlConfig{NewMemoryConfig(), url}
}

func (self *UrlConfig) Load() error {
	resp, err := http.Get(self.url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	out, err := unmarshalJson(body)
	if err != nil {
		return err
	}
	self.Reset(out)
	return nil
}

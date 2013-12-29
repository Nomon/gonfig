# Gonfig

Gonfig is a simple hierarchial config manager for Go lang

[![Build Status](https://travis-ci.org/Nomon/gonfig.png?branch=master)](https://travis-ci.org/Nomon/gonfig)
[![Coverage Status](https://coveralls.io/repos/Nomon/gonfig/badge.png?branch=HEAD)](https://coveralls.io/r/Nomon/gonfig?branch=HEAD)

### Api

All the config types including the root implement the Configurable interface:

```go
type Configurable interface {
  Get(string) interface{}
  Set(string, interface{})
  All() map[string]interface{}
  Reset(...map[string]interface{})
  Load() error
  Save() error
}
```

### Usage

```go
package main

import (
  . "github.com/Nomon/gonfig"
  "log"
)

func main() {
  conf := NewConfig()
  conf.Defaults().Reset(map[string]interface{}{
    "PATH": "/dome/configured/path",
    "a":    3711,
    "b":    2138,
    "c":    1908,
    "d":    912,
  })
  // use a configuration file
  conf.Use("json", NewJsonConfig("./config.json"))
  // also include all env variables, wioth optional prefix lookup
  // SETTINGS_DB="xyz" can be captured to "DB" with NewEnvConfig("SETTINGS_")
  conf.Use("env", NewEnvConfig(""))
  // Take in all commandline flags
  // prefix can be specified to pick only under named flagset
  // ie  NewArgvConfig("test")) to capture --test.asd into asd
  conf.Use("argv", NewArgvConfig(""))

  // real PATH in env, the lookup order for root config is the addition order,
  // so first form Defaults(), then from json, env, path and latest found is returned.
  log.Println("PATH in conf", conf.Get("PATH"))
  // /dome/configured/path
  // .Defaults() is shorthand for .Use("defaults")
  log.Println("Default PATH in conf", conf.Defaults().Get("PATH"))
  // /dome/configured/path
  log.Println("Default PATH in env conf", conf.Use("env").Get("PATH"))

  conf.Set("PATH", "/new/path")
  // the new changed path
  log.Println("PATH in conf", conf.Get("PATH"))
  // /dome/configured/path
  // .Set on root configuration wont override Defaults
  log.Println("Default PATH in conf", conf.Defaults().Get("PATH"))

  conf.Use("json").Set("abcd", "1234")
  if err := conf.Use("json").Save(); err != nil {
    log.Println(err)
    return
  }
  // reset config and reload from disk
  conf.Use("json").Reset()
  conf.Use("json").Load()
  // 1234
  log.Println("abcd from loaded json config", conf.Use("json").Get("abcd"))
}

```

Or Directly:

```go
func main() {
  conf := NewMomoryConfig();
  conf.Set("asd", "1234")
  log.Println("asd:", conf.Get("asd"))
  jsonconf := NewJsonConfig("./config.json");
  if err := jsonconf.Load(); err != nil {
    log.Println("Error loading conf", err)
    return
  }
  jsonconf.Set("asd", "1234")
  log.Println("asd:", jsonconf.Get("asd"))
  if err := jsonconf.Save(); err != nil {
    log.Println("config failed to save",err)
    return
  }
  log.Println("asd:1234 saved to config.json")
}
```

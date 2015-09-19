# Gonfig

Gonfig is a simple hierarchial config manager for Go lang. Inspired by [nconf](https://github.com/flatiron/nconf).

[![Build Status](https://travis-ci.org/Nomon/gonfig.png?branch=master)](https://travis-ci.org/Nomon/gonfig)
[![Coverage Status](https://coveralls.io/repos/Nomon/gonfig/badge.png?branch=HEAD)](https://coveralls.io/r/Nomon/gonfig?branch=HEAD)
## Docs

Available via: [![Go Walker](http://gowalker.org/api/v1/badge)](https://gowalker.org/github.com/Nomon/gonfig)

### Api

All the config types including the root implement the Configurable interface.

```go
// The main Configurable interface
// Also the hierarcial configuration (Config) implements it.
type Configurable interface {
  // Get a configuration variable from config
  Get(string) string
  // Set a variable, nil to reset key
  Set(string, string)
  // Reset the config data to passed data, if nothing is given set it to zero value
  Reset(...map[string]string)
  // Return a map of all variables
  All() map[string]string
}

type WritebleConfig interface {
  ReadableConfig
  // Save configuration
  Save() error
}

type ReadableConfig interface {
  Configurable
  // Load the configuration
  Load() error
}

// The hierarchial Config that can be used to mount other configs that are searched for keys by Get
type Config struct {
  // Overrides, these are checked before Configs are iterated for key
  Configurable
  // named configurables, these are iterated if key is not found in Config
  Configs map[string]Configurable
  // Defaults configurable, if key is not found in the Configurable & Configurables in Config,
  //Defaults is checked for fallback values
  Defaults Configurable
}

```

### Usage & Examples

For more examples check out [example_test.go](https://github.com/Nomon/gonfig/blob/master/example_test.go)


```go
  // Create a new root node for our hierarchical configuration
  conf := NewConfig(nil)

  // setting to the rootnode makes the variable found first, so it acts as an override for all the other
  // configurations set in the config.
  conf.Set("always", true);

  // use commandline variables myapp.*, ie --myapp-rules
  conf.Use("argv"), NewNewArgvConfig("myapp.*"))

  // use env variables MYAPP_*
  conf.Use("env", NewEnvConfig("MYAPP_"))

  // load local config file
  conf.Use("local", NewJsonConfig("./config.json"))

  // load global config file over the network
  conf.Use("global", NewUrlConfig("http://myapp.com/config/globals.json"))

  // Set some Defaults, if conf.Get() fails to find from any of the above configurations it will fall back to these.
  conf.Defaults.Reset(map[string]interface{}{
    "database": "127.0.0.1",
    "temp_dir":"/tmp",
    "always": false,
  })

  log.Println("Database host in network configuration",conf.Use("global").Get("database"))
  log.Println("Database host resolved from hierarchy",conf.Get("database"))

  // Save the hierarchy to a JSON file:
  jsonconf := NewJsonConf("./new_config.json")
  jsonconf.Reset(conf.All())
  if err := jsonconf.Save(); err != nil {
    log.Fatalln("Failed saving json config at path",jsonconf.Path,err)
  }
}

```

### Extending

Extending Gonfig is easy using the MemoryConfig or JsonConfig as a base, depending on the Save needs.
Here is an example implementation using a file with line separated key=value pairs for storage.


```go
type KVFileConfig struct {
  Configurable
  Path string
}

func NewKVFileConfig(path string) WritableConfig {
  return &KVFileConfig{NewMemoryConfig(), path}
}

func (self *KVFileConfig) Load() (err error) {
  var file *os.File

  if file, err = os.Open(self.Path); err != nil {
    return err
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    fmt.Println(scanner.Text())
    parts := strings.Split(scanner.Text(), "=")
    if len(parts) == 2 {
      self.Set(parts[0], parts[1])
    }
  }
  return scanner.Err();
}

func (self *KVFileConfig) Save() (err error) {
  var file *os.File;
  if file, err = os.Create(self.Path); err != nil {
    return err
  }
  defer file.Close();
  for k, v := range self.All() {
    if _, err = file.WriteString(k + "=" + fmt.Sprint("",v) + "\r\n"); err != nil {
      return err
    }
  }
  return nil
}
```


## License

```text
The MIT License (MIT)

Copyright (c) 2013 Matti Savolainen

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
```

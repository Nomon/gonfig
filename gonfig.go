package gonfig

// The hierarchial Config that can be used to mount other configs that are searched for keys by Get
type Config struct {
	configs map[string]Configurable
}

// The main Configurable interface
// All the single source configurations (NewJsonConfig, NewFileConfig, NewArgvConfig, NewEnvConfig, NewUrlConfig) implement it
// Also the hierarcial configuration (Config) implements it.
type Configurable interface {
	// Get a configuration variable from config
	Get(string) interface{}
	// Set a variable
	Set(string, interface{})
	// Return a map of all variables
	All() map[string]interface{}
	// Reset the config data to passed data, if nothing is given set it to zero value
	Reset(...map[string]interface{})
	// Loads the config, ie. from disk/url
	Load() error
	// Save the config. To file or Post to url.
	Save() error
}

// Creates a new config that is by default backed by a MemoryConfig Configurable
// Defaults can be accessed from .Defaults()
func NewConfig() *Config {
	cfg := &Config{
		configs: make(map[string]Configurable, 1),
	}
	cfg.Use("memory", NewMemoryConfig())
	cfg.Use("defaults", NewMemoryConfig())
	return cfg
}

// Returns the Defaults configuration that is used if no other config contains the desired key
// Returns the Defaults() memory configration
// This configuration is used if variable is not found in the hierarchy
// Defaults can be set to a configration:
//  conf.Defaults().Reset(map[string]interface{} (
//    "key": "value",
//  ))
func (self *Config) Defaults(config ...Configurable) Configurable {
	if len(config) > 0 {
		self.Use("defaults", config[0])
	}
	return self.Use("defaults")
}

// Resets all configs with the provided data, if no data is provided empties all stores
// Never touches the Defaults, to reset Defaults use config.Defaults().Reset()
func (self *Config) Reset(datas ...map[string]interface{}) {
	var data map[string]interface{}
	var store Configurable
	if len(datas) > 0 {
		data = datas[0]
	}
	if store = self.Defaults(); store == nil {
		store = NewMemoryConfig()
		self.Defaults(store)
	}
	for _, value := range self.configs {
		// dont reset defaults
		if value == store {
			continue
		}
		value.Reset(data)
	}
}

// Use config as named config and return an already set and loaded config
// mounts a new configuration in the hierarchy
// conf.Use("global", NewUrlConfig("http://host.com/config.json"))
// conf.Use("local", NewFileConfig("./config.json"))
// err := conf.Load();
// then get variable from specific config
// conf.Use("global").Get("key")
// or traverse the hierarchy and search for "key"
// conf.Get("key")
func (self *Config) Use(name string, config ...Configurable) Configurable {
	if len(config) == 0 {
		return self.configs[name]
	}
	self.configs[name] = config[0]
	self.configs[name].Load()
	return self.configs[name]
}

// Gets the key from first store that it is found from
func (self *Config) Get(key string) interface{} {
	defaults := self.Defaults()
	for _, config := range self.configs {
		if config == defaults {
			continue
		}
		if value := config.Get(key); value != nil {
			return value
		}
	}
	if value := defaults.Get(key); value != nil {
		return value
	}
	return nil
}

// Sets variable to all configurations except Defaults
func (self *Config) Set(key string, value interface{}) {
	for name, config := range self.configs {
		if name == "defaults" {
			continue
		}
		config.Set(key, value)
	}
}

// calls Load on all Configurables in Use
func (self *Config) Load() error {
	for _, config := range self.configs {
		if err := config.Load(); err != nil {
			return err
		}
	}
	return nil
}

// Saves all Configurables in use
func (self *Config) Save() error {
	if self.Defaults() == nil {
		self.Defaults(NewMemoryConfig())
	}
	for _, config := range self.configs {
		if err := config.Save(); err != nil {
			return err
		}
	}
	return nil
}

// Returns a map of data from all Configurables in use
// the first found instance of variable found is provided.
// cfg.Use("a", NewMemoryConfig())
// cfg.Use("b", NewMemoryConfig())
// cfg.Use("a").Set("a","1")
// cfg.Set("b").Set("a","2")
// then:
// cfg.All()["a"] == "1"
// cfg.Get("a") == "1"
// cfg.Use("b".).Get("a") == "2"
func (self *Config) All() map[string]interface{} {
	values := make(map[string]interface{}, 10)
	for _, config := range self.configs {
		for key, value := range config.All() {
			if values[key] == nil {
				values[key] = value
			}
		}
	}
	return values
}

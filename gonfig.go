// package gonfig provides tools for managing hierarcial configuration from multiple sources
package gonfig

// The hierarchial Config that can be used to mount other configs that are searched for keys by Get
type Config struct {
	Configurable
	// Defaults configurable, if key is not found in hierarchy Defaults will be checked.
	Defaults Configurable
	Configs  map[string]Configurable
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
	// Load/Save configuration
	Load() error
	Save() error
}

// Creates a new config that is by default backed by a MemoryConfig Configurable
// Defaults can be accessed from .Defaults()
func NewConfig(configs ...Configurable) *Config {
	cfg := &Config{
		NewMemoryConfig(),
		NewMemoryConfig(),
		make(map[string]Configurable),
	}
	return cfg
}

// Resets all configs with the provided data, if no data is provided empties all stores
// Never touches the Defaults, to reset Defaults use Config.Defaults().Reset()
func (self *Config) Reset(datas ...map[string]interface{}) {
	var data map[string]interface{}
	if len(datas) > 0 {
		data = datas[0]
	}
	for _, value := range self.Configs {
		if data != nil {
			value.Reset(data)
		} else {
			value.Reset()
		}
	}
	self.Configurable.Reset(data)
}

// Use config as named config and return an already set and loaded config
// mounts a new configuration in the hierarchy.
// conf.Use("global", NewUrlConfig("http://host.com/config..json")).
// conf.Use("local", NewFileConfig("./config.json"))
// err := conf.Load();.
// Then get variable from specific config.
// conf.Use("global").Get("key").
// or traverse the hierarchy and search for "key".
// conf.Get("key").
func (self *Config) Use(name string, config ...Configurable) Configurable {
	if self.Configs == nil {
		self.Configs = make(map[string]Configurable)
	}
	if len(config) == 0 {
		return self.Configs[name]
	}
	self.Configs[name] = config[0]
	self.Configs[name].Load()
	return self.Configs[name]
}

// Gets the key from first store that it is found from, checks Defaults
func (self *Config) Get(key string) interface{} {
	if value := self.Configurable.Get(key); value != nil {
		return value
	}

	for _, config := range self.Configs {
		if value := config.Get(key); value != nil {
			return value
		}
	}

	if value := self.Defaults.Get(key); value != nil {
		return value
	}
	return nil
}

// calls Configurable.Load() on all Configurable objects in the hierarchy.
func (self *Config) Load() error {
	for _, config := range self.Configs {
		if err := config.Load(); err != nil {
			return err
		}
	}
	if err := self.Configurable.Load(); err != nil {
		return err
	}
	return self.Defaults.Load()
}

// Saves all Configurables in use
func (self *Config) Save() error {
	for _, config := range self.Configs {
		if err := config.Save(); err != nil {
			return err
		}
	}
	return self.Configurable.Save()
}

// Returns a map of data from all Configurables in use
// the first found instance of variable found is provided.
// Config.Use("a", NewMemoryConfig()).
// Config.Use("b", NewMemoryConfig()).
// Config.Use("a").Set("a","1").
// Config.Set("b").Set("a","2").
// then.
// Config.All()["a"] == "1".
// Config.Get("a") == "1".
// Config.Use("b".).Get("a") == "2".
func (self *Config) All() map[string]interface{} {
	values := make(map[string]interface{})
	// put defaults in values
	for key, value := range self.Defaults.All() {
		if values[key] == nil {
			values[key] = value
		}
	}
	// put config values on top of them
	for _, config := range self.Configs {
		for key, value := range config.All() {
			if values[key] == nil {
				values[key] = value
			}
		}
	}
	// put overrides from self on top of all
	for key, value := range self.Configurable.All() {
		if values[key] == nil {
			values[key] = value
		}
	}
	return values
}

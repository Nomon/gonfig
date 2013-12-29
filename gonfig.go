package gonfig

type Config struct {
	configs map[string]Configurable
}

type Configurable interface {
	Get(string) interface{}
	Set(string, interface{})
	All() map[string]interface{}
	Reset(...map[string]interface{})
	Load() error
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
	for _, config := range self.configs {
		if value := config.Get(key); value != nil {
			return value
		}
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
	if self.Defaults() == nil {
		self.Defaults(NewMemoryConfig())
	}
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
func (self *Config) All() map[string]interface{} {
	values := make(map[string]interface{}, 10)
	for _, config := range self.configs {
		for key, value := range config.All() {
			values[key] = value
		}
	}
	return values
}

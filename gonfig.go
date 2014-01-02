// package gonfig provides tools for managing hierarcial configuration from multiple sources
package gonfig

// The main Configurable interface
// Also the hierarcial configuration (Config) implements it.
type Configurable interface {
	// Get a configuration variable from config
	Get(string) interface{}
	// Set a variable, nil to reset key
	Set(string, interface{})
	// Reset the config data to passed data, if nothing is given set it to zero value
	Reset(...map[string]interface{})
	// Return a map of all variables
	All() map[string]interface{}
}

type ReadableConfig interface {
	Configurable
	// Load the configuration
	Load() error
}

type WritableConfig interface {
	ReadableConfig
	// Save configuration
	Save() error
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

// Creates a new config that is by default backed by a MemoryConfig Configurable
// Takes optional initial configuration and an optional defaults
func NewConfig(initial Configurable, defaults ...Configurable) *Config {
	if initial == nil {
		initial = NewMemoryConfig()
	}
	def := NewMemoryConfig()
	if len(defaults) == 1 {
		def = defaults[0]
	}
	return &Config{
		initial,
		make(map[string]Configurable),
		def,
	}
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

	switch t := self.Configs[name].(type) {
	case ReadableConfig:
		t.Load()
	}
	return self.Configs[name]
}

// Gets the key from first store that it is found from, checks Defaults
func (self *Config) Get(key string) interface{} {
	// override from out values
	if value := self.Configurable.Get(key); value != nil {
		return value
	}
	// go through all in insert order untill key is found
	for _, config := range self.Configs {
		if value := config.Get(key); value != nil {
			return value
		}
	}
	// if not found check the defaults as fallback
	if value := self.Defaults.Get(key); value != nil {
		return value
	}

	return nil
}

func SaveConfig(config Configurable) error {
	switch t := config.(type) {
	case WritableConfig:
		if err := t.Save(); err != nil {
			return err
		}
	}
	return nil
}

// Saves all Configurables in use
func (self *Config) Save() error {
	for _, config := range self.Configs {
		if err := SaveConfig(config); err != nil {
			return err
		}
	}
	return SaveConfig(self.Configurable)
}

func LoadConfig(config Configurable) error {
	switch t := config.(type) {
	case ReadableConfig:
		if err := t.Load(); err != nil {
			return err
		}
	}
	return nil
}

// calls Configurable.Load() on all Configurable objects in the hierarchy.
func (self *Config) Load() error {
	LoadConfig(self.Configurable)
	LoadConfig(self.Defaults)
	for _, config := range self.Configs {
		LoadConfig(config)
	}
	return nil
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

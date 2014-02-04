package gonfig

// MemoryConfig is a simple abstraction to map[]interface{} for in process memory backed configuration
// only implements Configurable use JsonConfig to save/load if needed
type MemoryConfig struct {
	data map[string]string
}

// Returns a new memory backed Configurable
// The most basic Configurable simply backed by a map[string]interface{}
func NewMemoryConfig() *MemoryConfig {
	cfg := &MemoryConfig{make(map[string]string)}
	cfg.init()
	return cfg
}

func (self *MemoryConfig) init() {
	self.data = make(map[string]string)
}

// if no arguments are proced Reset() re-creates the underlaying map
func (self *MemoryConfig) Reset(datas ...map[string]string) {
	if len(datas) >= 1 {
		self.data = datas[0]
	} else {
		self.data = make(map[string]string)
	}
	return
}

// Get key from map
func (self *MemoryConfig) Get(key string) string {
	if self.data == nil {
		self.init()
	}
	return self.data[key]
}

// get all keys
func (self *MemoryConfig) All() map[string]string {
	if self.data == nil {
		self.init()
	}
	return self.data
}

// Set a key to value
func (self *MemoryConfig) Set(key, value string) {
	if self.data == nil {
		self.init()
	}
	self.data[key] = value
}

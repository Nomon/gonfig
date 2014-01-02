package gonfig

// MemoryConfig is a simple abstraction to map[]interface{} for in process memory backed configuration
type MemoryConfig struct {
	data map[string]interface{}
}

// Returns a new memory backed Configurable
// The most basic Configurable simply backed by a map[string]interface{}
func NewMemoryConfig() Configurable {
	cfg := &MemoryConfig{make(map[string]interface{})}
	cfg.init()
	return cfg
}

func (self *MemoryConfig) init() {
	self.data = make(map[string]interface{})
}

// if no arguments are proced Reset() re-creates the underlaying map
func (self *MemoryConfig) Reset(datas ...map[string]interface{}) {
	if len(datas) >= 1 {
		self.data = datas[0]
	} else {
		self.data = make(map[string]interface{})
	}
	return
}

func (self *MemoryConfig) Get(key string) interface{} {
	if self.data == nil {
		self.init()
	}
	return self.data[key]
}

func (self *MemoryConfig) All() map[string]interface{} {
	if self.data == nil {
		self.init()
	}
	return self.data
}

func (self *MemoryConfig) Set(key string, value interface{}) {
	if self.data == nil {
		self.init()
	}
	self.data[key] = value
}

func (self *MemoryConfig) Load() error {
	if self.data == nil {
		self.init()
	}
	return nil
}

func (self *MemoryConfig) Save() error {
	return nil
}

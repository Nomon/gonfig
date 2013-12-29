package gonfig

type memoryConfig struct {
	data map[string]interface{}
}

func NewMemoryConfig() Configurable {
	cfg := &memoryConfig{
		data: make(map[string]interface{}, 10),
	}
	cfg.Load()
	return cfg
}

func (self *memoryConfig) initialize() {
	self.data = make(map[string]interface{}, 10)
}

func (self *memoryConfig) Reset(datas ...map[string]interface{}) {
	if len(datas) == 0 {
		self.initialize()
		return
	}
	self.data = datas[0]
}

func (self *memoryConfig) Get(key string) interface{} {
	if self.data == nil {
		self.initialize()
	}
	return self.data[key]
}

func (self *memoryConfig) All() map[string]interface{} {
	if self.data == nil {
		self.initialize()
	}
	return self.data
}

func (self *memoryConfig) Set(key string, value interface{}) {
	if self.data == nil {
		self.initialize()
	}
	self.data[key] = value
}

func (self *memoryConfig) Load() error {
	if self.data == nil {
		self.initialize()
	}
	return nil
}

func (self *memoryConfig) Save() error {
	return nil
}

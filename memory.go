package gonfig

type MemoryConfig struct {
	data map[string]interface{}
}

func NewMemoryConfig() Configurable {
	cfg := &MemoryConfig{
		data: make(map[string]interface{}, 10),
	}
	cfg.Load()
	return cfg
}

func (self *MemoryConfig) initialize() {
	self.data = make(map[string]interface{}, 10)
}

func (self *MemoryConfig) Reset(datas ...map[string]interface{}) {
	if len(datas) == 0 {
		self.initialize()
		return
	}
	self.data = datas[0]
}

func (self *MemoryConfig) Get(key string) interface{} {
	if self.data == nil {
		self.initialize()
	}
	return self.data[key]
}

func (self *MemoryConfig) All() map[string]interface{} {
	if self.data == nil {
		self.initialize()
	}
	return self.data
}

func (self *MemoryConfig) Set(key string, value interface{}) {
	if self.data == nil {
		self.initialize()
	}
	self.data[key] = value
}

func (self *MemoryConfig) Load() error {
	if self.data == nil {
		self.initialize()
	}
	return nil
}

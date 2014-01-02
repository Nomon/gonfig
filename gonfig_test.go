package gonfig_test

import (
	. "github.com/Nomon/gonfig"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gonfig", func() {
	Describe("Config struct", func() {
		var cfg *Config
		BeforeEach(func() {
			cfg = NewConfig()
		})
		Describe("config.Default", func() {
			It("Should automatically create memory config for defaults", func() {
				defaults := cfg.Defaults
				Expect(defaults).ToNot(BeNil())
				memconf := NewMemoryConfig()
				memconf.Set("a", "b")
				cfg.Defaults.Reset(memconf.All())
				Expect(cfg.Get("a")).To(Equal("b"))
				Expect(cfg.Defaults.Get("a")).To(Equal("b"))
			})
		})
		It("Should use memory store to set and get by default", func() {
			cfg.Set("test_a", 10)
			Ω(cfg.Get("test_a")).Should(Equal(cfg.Get("test_a")))
		})
		It("Should return nil when key is non-existing", func() {
			Expect(cfg.Get("some-key")).To(BeNil())
		})
		It("Should return and use Defaults", func() {
			cfg.Defaults.Set("test_var", "abc")
			Ω(cfg.Defaults.Get("test_var")).Should(Equal("abc"))
			cfg.Set("test_var", "bca")
			Ω(cfg.Defaults.Get("test_var")).Should(Equal("abc"), "Setting to memory should not override defaults")
			Ω(cfg.Get("test_var")).Should(Equal("bca"), "Set to config should set in memory and use it over defaults")
		})

		It("Should reset everything else but Defaults() on reset", func() {
			cfg.Defaults.Set("test_var", "abc")
			Ω(cfg.Defaults.Get("test_var")).Should(Equal("abc"))
			cfg.Set("test_var", "bca")
			Ω(cfg.Defaults.Get("test_var")).Should(Equal("abc"), "Setting to memory should not override defaults")
			Ω(cfg.Get("test_var")).Should(Equal("bca"), "Set to config should set in memory and use it over defaults")
			cfg.Reset()
			Ω(cfg.Get("test_var")).Should(Equal("abc"), "Set to config should set in memory and use it over defaults")
		})

		It("Should load & save all relevant sources", func() {
			cfg.Use("json1", NewJsonConfig("./config_test_1.json"))
			cfg.Use("json2", NewJsonConfig("./config_test_2.json"))
			cfg.Use("json2").Set("asd", "123")
			cfg.Use("json1").Set("asd", "321")
			err := cfg.Save()
			Expect(err).ToNot(HaveOccurred())
			cfg.Reset()
			Expect(len(cfg.Use("json1").All()) == 0).To(BeTrue())
			err = cfg.Load()
			Expect(err).ToNot(HaveOccurred())
			Expect(cfg.Use("json1").Get("asd")).To(Equal("321"))
			Expect(cfg.Use("json2").Get("asd")).To(Equal("123"))
		})

		It("Should return all values from all storages", func() {
			cfg.Use("mem1", NewMemoryConfig())
			cfg.Use("mem2", NewMemoryConfig())
			cfg.Set("asd", "123456")
			cfg.Use("mem1").Set("das", "654321")
			cfg.Use("mem2").Set("sad", "654321")
			i := 0
			for key, value := range cfg.All() {
				Expect(cfg.Get(key)).To(Equal(value))
				i++
			}
			Expect(i == 3).To(BeTrue())
		})
		It("Should be able to use Config objects in the hierarchy", func() {
			cfg.Use("test", NewConfig())
			cfg.Set("test_123", "321test")
			Expect(cfg.Use("test").Get("test_123")).To(BeNil())
		})
		It("should prefere using defaults deeprer in hierarchy (reverse order to normal fetch.)", func() {
			deeper := NewConfig()
			deeper.Defaults.Reset(map[string]interface{}{
				"test":  123,
				"testb": 321,
			})
			cfg.Use("test", deeper)
			cfg.Defaults.Reset(map[string]interface{}{
				"test": 333,
			})
			Expect(cfg.Get("test")).To(Equal(123))
			Expect(cfg.Get("testb")).To(Equal(321))
			cfg.Set("testb", 1)
			Expect(cfg.Get("testb")).To(Equal(1))

		})
	})
})

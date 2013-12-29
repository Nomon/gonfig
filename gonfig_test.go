package gonfig_test

import (
	. "github.com/Nomon/gonfig"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gonfig", func() {
	Describe("Config", func() {
		var cfg *Config
		BeforeEach(func() {
			cfg = NewConfig()
		})
		It("Should use memory store to set and get by default", func() {
			cfg.Set("test_a", 10)
			Ω(cfg.Get("test_a")).Should(Equal(cfg.Get("test_a")))
		})
		It("Should return and use Defaults", func() {
			cfg.Defaults().Set("test_var", "abc")
			Ω(cfg.Defaults().Get("test_var")).Should(Equal("abc"))
			cfg.Set("test_var", "bca")
			Ω(cfg.Defaults().Get("test_var")).Should(Equal("abc"), "Setting to memory should not override defaults")
			Ω(cfg.Get("test_var")).Should(Equal("bca"), "Set to config should set in memory and use it over defaults")
		})
		It("Should reset everything else but Defaults() on reset", func() {
			cfg.Defaults().Set("test_var", "abc")
			Ω(cfg.Defaults().Get("test_var")).Should(Equal("abc"))
			cfg.Set("test_var", "bca")
			Ω(cfg.Defaults().Get("test_var")).Should(Equal("abc"), "Setting to memory should not override defaults")
			Ω(cfg.Get("test_var")).Should(Equal("bca"), "Set to config should set in memory and use it over defaults")
			cfg.Reset()
			Ω(cfg.Get("test_var")).Should(Equal("abc"), "Set to config should set in memory and use it over defaults")
		})
	})
})

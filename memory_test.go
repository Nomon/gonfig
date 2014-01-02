package gonfig_test

import (
	. "github.com/Nomon/gonfig"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MemoryConfig", func() {
	var cfg Configurable
	BeforeEach(func() {
		cfg = NewMemoryConfig()
	})
	It("Should be able to store and retrieve data", func() {
		cfg.Set("test", "abc")
		Expect(cfg.Get("test")).To(Equal("abc"))
	})
	It("Should Reset() to zero length All()", func() {
		cfg.Set("test", "abc")
		Expect(cfg.Get("test")).To(Equal("abc"))
		Expect(len(cfg.All())).To(Equal(1))
		cfg.Reset()
		Expect(len(cfg.All())).To(Equal(0))
	})
	It("Should Reset() with data correctly", func() {
		cfg.Set("test", "abc")
		cfg.Set("test_abc", "123")
		Expect(cfg.Get("test")).To(Equal("abc"))
		Expect(len(cfg.All())).To(Equal(2))
		cfg2 := NewMemoryConfig()
		cfg2.Reset(cfg.All())
		Expect(cfg.Get("test")).To(Equal(cfg2.Get("test")))
	})
})

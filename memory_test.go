package gonfig_test

import (
	. "github.com/nomon/gonfig"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MemoryConfig", func() {
	It("Should be able to store and retrieve data", func() {
		cfg := NewMemoryConfig()
		cfg.Set("test", "abc")
		Ω(cfg.Get("test")).Should(Equal("abc"))
	})
})

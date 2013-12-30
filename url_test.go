package gonfig_test

import (
	"fmt"
	. "github.com/nomon/gonfig"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UrlConfig", func() {
	var (
		cfg *Config
		err error
	)
	BeforeEach(func() {
		cfg = NewConfig()
	})
	JustBeforeEach(func() {
		cfg.Use("url", NewUrlConfig(fmt.Sprintf("http://127.0.0.1:%d", HttpPort)))
	})
	It("Should load config from URL", func() {
		Expect(cfg).ToNot(BeNil())
		Expect(err).ToNot(HaveOccurred())
		err := cfg.Load()
		Expect(err).ToNot(HaveOccurred())
		Expect(cfg.Get("test")).To(Equal("abc"))
	})
})

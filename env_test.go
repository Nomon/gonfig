package gonfig_test

import (
	. "github.com/nomon/gonfig"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EnvConfig", func() {
	var (
		err error
		cfg Configurable
	)
	BeforeEach(func() {
		cfg = NewEnvConfig("")
		err = cfg.Load()
	})
	It("Should load variables from environment", func() {
		Expect(len(cfg.All()) > 0).To(BeTrue())
	})
})

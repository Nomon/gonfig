package gonfig_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
	"testing"
)

func TestGonfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gonfig Suite")
	os.Remove("./config_test_1.json")
	os.Remove("./config_test_2.json")
	os.Remove("./config.json")
}

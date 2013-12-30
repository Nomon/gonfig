package gonfig_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/rand"
	"net/http"
	"os"
	"testing"
	"time"
)

var (
	HttpPort int
)

func randPort(start int, end int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return start + rand.Intn(end-start)
}
func dummyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"test":"abc","test_b":123}`)
}

func TestGonfig(t *testing.T) {
	// start test http server to serve dummy json
	HttpPort = randPort(1024, 30000)
	http.HandleFunc("/", dummyHandler)
	go func() {
		http.ListenAndServe(fmt.Sprintf(":%d", HttpPort), nil)
	}()

	RegisterFailHandler(Fail)
	RunSpecs(t, "Gonfig Suite")
	os.Remove("./config_test_1.json")
	os.Remove("./config_test_2.json")
	os.Remove("./config.json")
}

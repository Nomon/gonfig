package gonfig

import (
	"flag"
	"log"
	"os"
	"strings"
)

// the Argv configurable.
type ArgvConfig struct {
	Configurable
	Prefix string
}

// Creates a new ArgvConfig and returns it as a ReadableConfig
func NewArgvConfig(prefix string) ReadableConfig {
	cfg := &ArgvConfig{NewMemoryConfig(), prefix}
	return cfg
}

// Loads all the variables from argv to the underlaying Configurable.
// If a Prefix is provided for ArgvConfig then keys are imported with the Prefix removed
// so --test.asd=1 with Prefix 'test.' imports "asd" with value of 1
func (self *ArgvConfig) Load() (err error) {
	flagset := flag.NewFlagSet("arguments", flag.ContinueOnError)
	flagset.Parse(os.Args)

	flagset.VisitAll(func(f *flag.Flag) {
		name := f.Name
		log.Println(f.Name)
		if self.Prefix != "" && strings.HasPrefix(f.Name, self.Prefix) {
			name = strings.Replace(name, self.Prefix, "", 1)
		}
		if getter, ok := f.Value.(flag.Getter); ok {
			self.Set(name, getter.Get().(string))
		}
	})
	return nil
}

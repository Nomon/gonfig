package gonfig_test

import (
	"fmt"
	. "github.com/Nomon/gonfig"
	"os"
)

func ExampleHierarchy() {
	conf := NewConfig(nil)             // root config
	conf.Use("second", NewConfig(nil)) // config in hierarchy as second
	conf.Use("second").Set("asd", "abc")
	fmt.Println(conf.Get("asd"))
	// Output: abc
}

func ExampleDefaults() {
	conf := NewConfig(nil) // root config
	conf.Defaults.Reset(map[string]string{
		"test_default":   "123",
		"test_default_b": "321",
	})
	conf.Use("second", NewConfig(nil)) // config in hierarchy as second
	conf.Use("second").Set("test_default", "333")
	fmt.Println(conf.Get("test_default"), conf.Get("test_default_b"))
	// Output: 333 321
}

func ExampleSaveToJson() {
	conf := NewConfig(nil)
	conf.Set("some", "variable")
	jsonconf := NewJsonConfig("./config.json", conf)
	jsonconf.Save()
	// OR:
	// jsonconf := NewJsonConf("./config")
	// jsonconf.Reset(conf.All());
	// jsonconf.Save()

	if err := jsonconf.Save(); err != nil {
		fmt.Println("Error saving config", err)
	}
	jsonconf2 := NewJsonConfig("./config.json")
	if err := jsonconf2.Load(); err != nil {
		fmt.Println("Error loading config", err)
	}
	fmt.Println(jsonconf2.Get("some"))
	os.Remove("./config.json")
	// Output: variable
}

func ExampleConstruction() {
	var cfg MemoryConfig
	cfg.Set("example1", "123")
	fmt.Println(cfg.Get("example1"))
	cfg2 := NewConfig(&cfg)
	fmt.Println(cfg2.All())
	// Output: 123
}

package gonfig_test

import (
	"fmt"
	. "github.com/Nomon/gonfig"
	"os"
)

func ExampleHierarchy() {
	conf := NewConfig()             // root config
	conf.Use("second", NewConfig()) // config in hierarchy as second
	conf.Use("second").Set("asd", "abc")
	fmt.Println(conf.Get("asd"))
	// Output: abc
}

func ExampleDefaults() {
	conf := NewConfig() // root config
	conf.Defaults().Reset(map[string]interface{}{
		"test_default":   123,
		"test_default_b": 321,
	})
	conf.Use("second", NewConfig()) // config in hierarchy as second
	conf.Use("second").Set("test_default", 333)
	fmt.Println(conf.Get("test_default"), conf.Get("test_default_b"))
	// Output: 333 321
}

func ExampleSaveToJson() {
	conf := NewConfig()
	conf.Set("some", "variable")
	jsonconf := NewJsonConfig("./config.json")
	jsonconf.Reset(conf.All())
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
package gonfig_test

import (
	"fmt"
	. "github.com/Nomon/gonfig"
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

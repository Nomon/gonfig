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

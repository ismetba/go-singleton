package main

import (
	"fmt"

	"github.com/ismetbayandur/singletonTest/singleton"
	"github.com/ismetbayandur/singletonTest/test"
)

func main() {
	s := singleton.Get()
	s.Name = "ismet"
	fmt.Println(s)
	test.Test()
	fmt.Println(s)
}

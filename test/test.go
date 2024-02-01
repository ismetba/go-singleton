package test

import "github.com/ismetbayandur/singletonTest/singleton"

func Test() {
	s := singleton.Get()
	s.Surname = "bayandur"
}

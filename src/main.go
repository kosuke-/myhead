package main

import (
	"flag"
)

var (
	number = flag.Int("n", 10, "head line number")
)

func main() {
	flag.Parse()
	m := NewMyhead(*number,flag.Args())
	m.Run()
}

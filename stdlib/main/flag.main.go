package main

import (
	"flag"
	"fmt"
)

var (
	name string
	age  int
)

func main() {
	fmt.Printf("name: %s, age: %d\n", name, age)
}

func init() {
	flag.StringVar(&name, "name", "default-name", "help message for name")
	flag.IntVar(&age, "age", 18, "help message for name")
	flag.Parse()
}

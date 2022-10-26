package main

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	LoadConfig(".")
	x := Cfg("urls.shopping")

	fmt.Println("xxxx", x)
}

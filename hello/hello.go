package main

import (
	"fmt"
	"github.com/golang/golearning/golib"
	)

func main() {

	fmt.Printf("Hello Go!\n")

	var t int;

	t = golib.Add(10,22)

	fmt.Printf("Sume of 10, and 22 is: %d\n", t)

}

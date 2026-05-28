package main

import (
	"fmt"
)

var a int
var b float64
var c string
var d bool

func main() {
	fmt.Printf("%v, var de 'a' é %T\n", a, a)
	fmt.Printf("%v, var de 'b' é %T\n", b, b)
	fmt.Printf("%v, var de 'c' é %T\n", c, c)
	fmt.Printf("%v, var de 'd' é %T\n", d, d)
}

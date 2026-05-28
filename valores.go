package main

import (
	"fmt"
)

var a int     // Declaração de variavel para valor int
var b float64 // Declaração de variavel para valor float64
var c string  // Declaração de variavel para valor string
var d bool    // Declaração de variavel para valor bool

func main() { //print para variaveis sem valor atribuido
	fmt.Printf("%v, var de 'a' é %T\n", a, a)
	fmt.Printf("%v, var de 'b' é %T\n", b, b)
	fmt.Printf("%v, var de 'c' é %T\n", c, c)
	fmt.Printf("%v, var de 'd' é %T\n", d, d)
}

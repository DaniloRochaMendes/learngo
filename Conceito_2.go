package main

import "fmt"

var tst1 = "Print Fora da função main" // Declaração de variável global

func main() { //Abertura da função main
	gopher := "Declaração da Variavel Gopher. O nome da declaração ':='"
	fmt.Println(gopher)
	fmt.Print("tst1: ", tst1, " com ", gopher, "\n") // Gera a variavel global

	tst1 = "\n Print Dentro da função main \n" // Atribuição de variável local com o mesmo nome da variável global
	tst2 := 20                                 // declaração e atribuição de variavel int
	fmt.Printf(tst1)                           // Gera a variável local
	fmt.Printf("tst2: %v, %T\n", tst2, tst2)

}

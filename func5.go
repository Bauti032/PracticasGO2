package main

import (
	"fmt"
	"math/rand"
)


func main() {

	var (
		numero_aleatorio int
	)


	numero_aleatorio = generarNumero()

}


func generarNumero() int {
	return rand.Intn(101) // Funcion que genera el numero aleatorio
}

func checknumeroadivinado (numero_aleatorio int) int {
	
}

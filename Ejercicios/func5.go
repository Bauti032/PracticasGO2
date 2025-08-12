package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	numero_aleatorio := generarNumero()
	var intentos int
	var adivinado bool

	for !adivinado {
		var numero_usuario int
		fmt.Print("Introduce un número entre 0 y 100: ")
		fmt.Scan(&numero_usuario)
		intentos++

		resultado := checkNumeroAdivinado(numero_usuario, numero_aleatorio)
		procesarResultado(resultado, intentos)

		if resultado == 0 {
			adivinado = true
			calificarJugador(intentos)
		}
	}
}

func generarNumero() int {
	return rand.Intn(101) // Generar numero aleatorio entre 0  y 100
}

// Devuelve -1 si usuario < aleatorio, 1 si usuario > aleatorio, 0 si son iguales
func checkNumeroAdivinado(numero_usuario int, numero_aleatorio int) int {
	if numero_usuario < numero_aleatorio {
		return -1
	} else if numero_usuario > numero_aleatorio {
		return 1
	}
	return 0
}

// Imprimir si numero es menor, mayor o igual
func procesarResultado(resultado int, intentos int) {
	switch resultado {
	case -1:
		fmt.Printf("El número a adivinar es mayor. Intentos: %d\n", intentos)
	case 1:
		fmt.Printf("El número a adivinar es menor. Intentos: %d\n", intentos)
	case 0:
		fmt.Printf("¡Correcto! Has adivinado el número en %d intentos.\n", intentos)
	}
}

// Calificar jugador
func calificarJugador(intentos int) {
	if intentos <= 5 {
		fmt.Println("¡Eres bueno!")
	} else if intentos < 15 {
		fmt.Println("Eres regular.")
	} else {
		fmt.Println("No eres muy bueno.")
	}
}


// El procedimiento no devuelve resultados. Esa es la diferencia con la funcion y el procedimieto 
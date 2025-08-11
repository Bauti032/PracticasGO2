package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		A         int
		B         int
		operacion int
		op        int
	)

	for {
		fmt.Print("Ingresar que tipo de operacion desea hacer (1 suma, 2 resta, 3 multiplicacion, 4 division, 0 salir del programa): ")
		fmt.Scanln(&operacion)

		if operacion == 0 {
			salir()
		}

		fmt.Print("Ingresar los dos numeros a operar: ")
		fmt.Scanln(&A)
		fmt.Scanln(&B)

		switch operacion {
		case 1:
			op = suma(A, B)
		case 2:
			op = resta(A, B)
		case 3:
			op = multiplicacion(A, B)
		case 4:
			op = division(A, B)
		default:
			fmt.Println("Operacion no valida")
			continue
		}
		fmt.Println("El resultado de la operacion es: ", op)
	}
}

func suma(A int, B int) int {
	return A + B
}

func resta(A int, B int) int {
	return A - B
}

func multiplicacion(A int, B int) int {
	return A * B
}

func division(A int, B int) int {
	if B == 0 {
		fmt.Println("Error: division por cero")
		return 0
	}
	return A / B
}

func salir() {
	os.Exit(0)
}

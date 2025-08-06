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
	)

	fmt.Print("Ingresar que tipo de operacion desea hacer (1 suma, 2 resta, 3 multiplicacion, 4 division, 0 salir del programa): ")
	fmt.Scanln(&operacion)
	fmt.Print("Ingresar los dos numeros a operar: ")
	fmt.Scanln(&A)
	fmt.Scanln(&B)

		for operacion != 0{
			switch operacion {
			case 1:
				op := suma(operacion, A, B)
			case 2:
				op := resta(operacion, A, B)
			case 3:
				op := multiplicacion(operacion, A, B)
			case 4:
				op := division(operacion, A, B)

			case 0:
				op := salir(operacion)
			}
		}
		fmt.Print("El resultado de la operacion es: ", op)

	
}

func suma(operacion int, A int, B int) int {
	return A + B
}

func resta(operacion int, A int, B int) int {
	return A - B

}

func multiplicacion(operacion int, A int, B int) int {
	return A * B

}

func division(operacion int, A int, B int) int {
	return A / B

}

func salir(operacion int) int {
	os.Exit(0)
	return operacion

}

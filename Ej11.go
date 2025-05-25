package main

import ("fmt"
		"math"
)
 
var (
			N1 int
			R int
			Pares int
			Impares int
		)
func main() {
	
	for x := 1; x <= 10; x++ {
		fmt.Println("Ingrese el numero: ")
		fmt.Scanln(&N1)
		R = int(math.Pow(-1, float64(N1)))
	
		if R > 0 { 
			Pares =+ N1
	} else {
			Impares =+ N1
		}
	
	}
	fmt.Println("La suma de los numeros pares es: ", Pares)
	fmt.Println("La suma de los numeros impares es: ", Impares) 

	}
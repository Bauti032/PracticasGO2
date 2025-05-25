package main
import (
	"fmt"
)
var (
	Suma int
	N1 int
)

func main () {
	fmt.Println("Ingrese el Numero a sumar: ")
	fmt.Scanln(&N1)

	for x := 0; x < N1; x++ {
		Suma += x
		
	} 
	fmt.Printf("La suma del numero es %d", Suma)
}
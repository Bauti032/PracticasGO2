package main
import (
	"fmt"
)
var (
	N1 int
	NDistinto int
	Mayor int
	Menor int
)
func main () {
	for x:= 0; x <= 9; x ++ {
		fmt.Println("Ingrese el numero: ")
		fmt.Scanln(&N1)

		if x == 1 {
			NDistinto = N1
			Mayor = N1
			Menor = N1
		} 
		if N1 > Mayor {
			Mayor = N1
		}
		if N1 < Menor {
			Menor = N1
		}
		
	}
	fmt.Printf("El mayor es %d, y el menor es %d", Mayor, Menor)
}

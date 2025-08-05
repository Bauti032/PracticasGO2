package main
import (
	"fmt"
)

func main() {

var (
	Cdolares int
	CPesos int
	CotDia int 
)

	fmt.Println("Ingrese la cotización del dia: ")
	fmt.Scanln(&CotDia)

	fmt.Println("Ingrese la cantidad de Pesos: ")
	fmt.Scanln(&CPesos)

	for CPesos != 0 {
		fmt.Println("Ingrese la cantidad de dolares: ")
		fmt.Scanln(&Cdolares)

		Resultado := Operacion(CotDia, CPesos, Cdolares)

		fmt.Println(Resultado)

		fmt.Println("Ingrese la cantidad de pesos: ")
		fmt.Scanln(&CPesos)
	}

}

func Operacion(CotDia int, CPesos int, Cdolares int) bool {

	Resultado := CPesos / CotDia

	if Resultado >= Cdolares {
		// fmt.Println("Transacción correcta")
		return true
	} else {
		// fmt.Println("No se ah podido hacer la transacción por fondos insuficientes")
		return false
	}
	

}
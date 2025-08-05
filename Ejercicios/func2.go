package main

import (
	"fmt"
)

func main(){
var (
	numero int
	cn     int = 5
)

for x:= 0; x < cn; x++{
	fmt.Print("Ingrese un numero: ")
	fmt.Scanln(&numero)
	
	res1 := verificarMultiplo(numero, 3)
	res2 := verificarMultiplo(numero, 5)

	if res1 && !res2 {
		fmt.Println("El numero es multiplo de 3 y no de 5")
	}else if !res1 && res2 {
		fmt.Println("El numero es multiplo de 5 y no de 3")
	}
}


}

func verificarMultiplo(numero int, m int) bool {

	if numero % m == 0 {
		return true
	}else {
		return false
	}
}
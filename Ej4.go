package main

import "fmt"

func main() {
	var Numero int

	fmt.Println("Ingrese el Numero del 1 al 7: ")
	fmt.Scanln(&Numero)

	if Numero >= 1 && Numero <= 7 {

		switch Numero {
		case 1:
			fmt.Println("Es lunes")
		case 2:
			fmt.Println("Es martes")
		case 3:
			fmt.Println("Es miercoles")
		case 4:
			fmt.Println("Es jueves")
		case 5:
			fmt.Println("Es viernes")
		case 6:
			fmt.Println("Es sabado")
		case 7:
			fmt.Println("Es domingo")
		}

	} else {
		fmt.Print("El numero ingresado es incorrecto.  ")
		return
	}

}

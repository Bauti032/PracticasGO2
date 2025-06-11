package main

import (
	"fmt"
)

func main() {
	var (
		n   string
		e   int
		s   int
		nem string
		ei  int
		acs int
		m   int

	)


		fmt.Print("Ingrese la edad del empleado: ")
		fmt.Scanln(&e)

	for e != 0 {
		fmt.Print("Ingrese el nombre del empleado: ")
		fmt.Scanln(&n)
		fmt.Print("Ingrese el sueldo del empleado: ")
		fmt.Scanln(&s)

		if e > 25 && s > 3000 {
			nem = nem + n

		}

		if e < 18 {
			m++
		}

		ei = ei + 1
		acs += s
		
		fmt.Print("Ingrese la edad del empleado: ")
		fmt.Scanln(&e)
	}

	fmt.Println("Empleados mayores de 25 años con sueldo mayor a 3000:", nem)
	fmt.Println("Cantidad de empleados menores de 18 años:", m)
	fmt.Println("Cantidad de empleados ingresados:", ei)
	fmt.Println("Sueldo total de los empleados:", acs)

}

package main

import (
	"fmt"
)

const (
	cantcolumnas = 6
)

func main() {

	var (
		datos [][cantcolumnas]int
	)

	matriz := cargamat(datos)

	fmt.Println(matriz)
}

func cargamat(datos [][cantcolumnas]int) [][cantcolumnas]int {

	var (
		vec         [6]int
		codcli      int
		codanimal   int
		codservicio int
		importe     int
		atdom       int
	)

	nro_serv := 999

	for {
		fmt.Println("Ingrese el codigo de cliente: ")
		fmt.Scanln(&codcli)

		if codcli == 0 {
			break
		}

		fmt.Println("Ingrese el codigo clase de animal: ")
		fmt.Scanln(&codanimal)
		fmt.Println("Ingrese el codigo de servicio: ")
		fmt.Scanln(&codservicio)
		fmt.Scanln("Ingrese el importe total: ")
		fmt.Scanln(&importe)
		fmt.Scanln("Ingrese si tuvo atencion a domicilio (1 si - 2 no): ")
		fmt.Scanln(&atdom)
		nro_serv = +1
		vec[0] = nro_serv
		vec[1] = codcli
		vec[2] = codanimal
		vec[3] = codservicio
		vec[4] = importe
		vec[5] = atdom

		datos = append(datos, vec)
	}

	return datos
}

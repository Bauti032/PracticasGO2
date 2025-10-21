package main

import "fmt"

const (
	coltel     = 2
	colllamada = 3
)

func main() {

	var (
		mattelef [][coltel]int
		matllam  [][colllamada]int
	)

	cargatel(&mattelef)
	cargallam(mattelef, &matllam)

	e1 := totalminutos(mattelef, matllam)
	fmt.Println("El total de minutos en llamada por usuarios es: ", e1)
}

func cargatel(mattelef *[][coltel]int) {

	var (
		num     int
		empresa int
		vec     [2]int
	)

	for {
		fmt.Println("Ingrese el Numero de telefono (0 para salir): ")
		fmt.Scanln(&num)

		if num == 0 {
			break
		}

		fmt.Println("Ingrese la empresa (1- Claro, 2-Movistar, 3-Personal): ")
		fmt.Scanln(&empresa)

		vec[0] = num
		vec[1] = empresa

		*mattelef = append(*mattelef, vec)

	}

}
func cargallam(mattelef [][coltel]int, matllam *[][colllamada]int) {

	var (
		norigen int
		ndest   int
		min     int
		vec     [3]int
	)

	fmt.Println("Numeros disponibles: ")
	for y := 0; y < len(mattelef); y++ {
		fmt.Printf("%d. Número: %d \n", y, mattelef[y][0])
	}

	for {

		fmt.Println("Ingrese el Numero de Origen: ")
		fmt.Scanln(&norigen)

		if norigen == 0 {
			break
		}

		fmt.Println("Ingrese el Numero de destino: ")
		fmt.Scanln(&ndest)
		fmt.Println("Ingrese la cantidad de minutos: ")
		fmt.Scanln(&min)

		vec[0] = norigen
		vec[1] = ndest
		vec[2] = min

		*matllam = append(*matllam, vec)

	}

	fmt.Print(matllam)

}

/*Realizar una función que devuelva como resultado el total de minutos consumidos de todos los
usuarios de una determinada empresa pasada como parámetro. Imprimir el resultado
devuelto. Parámetros: 3.*/

func totalminutos(mattelef [][coltel]int, matllam [][colllamada]int) int {
	var (
		acm1    int
		empresa int
	)

	fmt.Print("Ingrese la empresa: ")
	fmt.Scanln(&empresa)

	for f := 0; f < len(mattelef); f++ {
		for f1 := 0; f1 < len(matllam); f1++ {
			if matllam[f1][0] == mattelef[f][0] && mattelef[f][1] == empresa {
				acm1 += matllam[f1][2]
			}
		}
	}

	return acm1
}

/*d) Realizar una función que devuelva la cantidad de llamadas no contestadas por un número
destino pasado como parámetro (si cantidad total de minutos de llamada es igual a cero, es
porque el destino no atendió la llamada). Imprimir el resultado devuelto. Parámetros: 2.*/

func cantllamadas(matllam [][colllamada]int) {
	var ()

}

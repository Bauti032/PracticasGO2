package main

import (
	"fmt"
)

func main() {
	var (
		mat  int
		par1 int
		par2 int
		par3 int
		ca   int
		cap  int
		pa   int
	)

	fmt.Print("Ingrese el numero de matricula: ")
	fmt.Scanln(&mat)

	for mat != 0 {
		ca++

		fmt.Print("Ingrese nota del parcial 1: ")
		fmt.Scanln(&par1)

		fmt.Print("Ingrese nota del parcial 2: ")
		fmt.Scanln(&par2)

		fmt.Print("Ingrese nota del parcial 3: ")
		fmt.Scanln(&par3)

		pa = 0
		if par1 >= 7 {
			pa++
		} else if par2 >= 7 {
			pa++
		} else if par3 >= 7 {
			pa++
		}

		if pa >= 2 {
			cap++
		}

		fmt.Print("Ingrese el numero de matricula: ")
		fmt.Scanln(&mat)
	}

	porcentaje := (ca / cap) * 100

	fmt.Printf("Cantidad total de alumnos %d, cantidad de alumnos regularizados %d, porcentaje de aprobados %d.", ca, cap, porcentaje)
}

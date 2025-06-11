// N votantes, sexo H o M, candidatos del 1 al 3. Cantidad de mujeres que votaron, cantidad de varones que votaron al C3. Candidato ganador sin empate
package main

import (
	"fmt"
)

var (
	Sex_Option       string
	Candidate_number int
	Qwomen           int
	Qmen3            int
	Candidate_win    int
	Vq               int
	Candidate1       int
	Candidate2       int
	Candidate3       int
)

func main() {
	fmt.Println("Ingrese la cantidad de votantes: ")
	fmt.Scanln(&Vq)
	for x := 0; x < Vq; x++ {
		fmt.Println("Ingrese el sexo: ")
		fmt.Scanln(&Sex_Option)
		fmt.Println("Ingrese el N° de candidato: ")
		fmt.Scanln(&Candidate_number)

		if Candidate_number == 1 {
			Candidate1++
			if Sex_Option == "M" {
				Qwomen++
			}
		}
		if Candidate_number == 2 {
			Candidate2++
			if Sex_Option == "M" {
				Qwomen++
			}
		}
		if Candidate_number == 3 {
			Candidate3++
			if Sex_Option == "H" {
				Qmen3++
			} else {
				Qwomen++
			}
		}
	}
	if Candidate1 > Candidate2 && Candidate1 > Candidate3 {
		Candidate_win = 1
	} else if Candidate2 > Candidate1 && Candidate2 > Candidate3 {
		Candidate_win = 2
	} else  {
		Candidate_win = 3
	}

	fmt.Printf("El candidato ganador es %d. La cantidad de mujeres que votaron son %d, la cantidad de hombres que votaron al C3 es %d", Candidate_win, Qwomen, Qmen3)

}

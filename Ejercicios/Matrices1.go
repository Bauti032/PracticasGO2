package main

import "fmt"

func main() {
	var (
		A[4][3] int
	)

	Res := Suma(A)
	fmt.Println(Res)
	
}

func Suma(A[4][3]int) [2]int {
	var n int
	var cintervino int
	var cmayor int
	var R[2]int

	for x := 0; x < 4; x++ {
		for y := 0; y < 3; y++ {
			fmt.Println("Ingrese el Número: ")
			fmt.Scanln(&n)
			A[x][y] = n
		}
	}

	for x := 0; x < 4; x++ {
		for y := 0; y < 3; y++ {
			if A[x][y] >= 5 {
				cmayor += 1
				cintervino += 1
			} else {
				cintervino += 1
			}
		}
	}
	R[0] = cmayor
	R[1] = cintervino
	return R
}

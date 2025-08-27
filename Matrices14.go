package main

import "fmt"

// 14) Cargar una matriz A de 5x3 y un vector B de 10 elementos. Sumar los elementos de A, que no encuentren su
// igual en B. Imprimir la suma obtenida.

const(
	f int = 5
	c int = 3
	EB int = 10
)

func main(){
	var (
		A [f][c]int
		B [EB]int
	)

	CargarMyV(&A, &B)
	Cigualdad(A, B)

	fmt.Print(ElementosNI)
	 
}

func CargarMyV(A *[f][c]int, B *[EB]int){


for x:=0; x<f; x++{
	for y:=0; y<c; y++{
		fmt.Printf("Ingrese un valor para fila %d, columna %d: ", x+1, y+1)
		fmt.Scanln(&A[x][y])
	}
}


for x:=0; x<EB; x++{
	fmt.Printf("Ingrese un valor para el elemento %d: ", x+1)
	fmt.Scanln(&B[x])
}


}


func Cigualdad(A [f][c]int, B [EB]int)int{
	var(
		ElementosNI int
	)


			for x:=0; x<f; x++{
				for y:=0; y<c; y++{
					for z:=0; z<EB; z++{
						if A[x][y] == B[z]{
							A[x][y] = 0
						}else{
							ElementosNI =+ A[x][y] 
						}
					}
				}
			}

			return ElementosNI
}

